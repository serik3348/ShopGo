package entities

type User struct {
	Id             int64
	Nameandsurname string `validate:"required" label:"Name and Surname" json:"nameandsurname"`
	Email          string `validate:"required,email,isunique=users-email"json:"email"`
	Username       string `validate:"required,gte=3,isunique=users-username"json:"username"`
	Password       string `validate:"required,gte=6"json:"password"`
	Cpassword      string `validate:"required,eqfield=Password" label:"Confirmation Password"json:"cpassword"`
}
