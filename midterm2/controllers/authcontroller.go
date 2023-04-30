package controllers

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"server/config"
	"server/entities"
	"server/libraries"
	"server/models"
)

type UserInput struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

var userModel = models.NewUserModel()
var validation = libraries.NewValidation()

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			data := map[string]interface{}{
				"nameandsurname": session.Values["nameandsurname"],
			}

			temp, _ := template.ParseFiles("midterm2/views/index.html")
			temp.Execute(w, data)
		}
	}

}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("midterm2/views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		UserInput := UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		errorMessages := validation.Struct(UserInput)
		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
			}
			temp, _ := template.ParseFiles("midterm2/views/login.html")
			temp.Execute(w, data)
		} else {
			var user entities.User
			userModel.Where(&user, "username", UserInput.Username)

			var message error
			if user.Username == "" {
				message = errors.New("Incorrect Username or Password")
			} else {
				errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
				if errPassword != nil {
					message = errors.New("Incorrect Username or Password")
				}

			}
			if message != nil {
				data := map[string]interface{}{
					"error": message,
				}
				temp, _ := template.ParseFiles("midterm2/views/login.html")
				temp.Execute(w, data)
			} else {
				session, _ := config.Store.Get(r, config.SESSION_ID)

				session.Values["loggedIn"] = true
				session.Values["email"] = user.Email
				session.Values["username"] = user.Username
				session.Values["nameandsurname"] = user.Nameandsurname

				session.Save(r, w)

				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("midterm2/views/register.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()

		user := entities.User{
			Nameandsurname: r.Form.Get("nameandsurname"),
			Email:          r.Form.Get("email"),
			Username:       r.Form.Get("username"),
			Password:       r.Form.Get("password"),
			Cpassword:      r.Form.Get("cpassword"),
		}

		errorMessages := validation.Struct(user)
		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
				"user":       user,
			}
			temp, _ := template.ParseFiles("midterm2/views/register.html")
			temp.Execute(w, data)
		} else {
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			userModel.Create(user)

			data := map[string]interface{}{
				"pesan": "Successful registration",
			}
			temp, _ := template.ParseFiles("midterm2/views/register.html")
			temp.Execute(w, data)
		}
	}
}
