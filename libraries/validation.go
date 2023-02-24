package libraries

import (
	"database/sql"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trasnlations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"server/config"
	"strings"
)

type Validation struct {
	conn *sql.DB
}

func NewValidation() *Validation {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}
	return &Validation{conn: conn}
}

func (v *Validation) Init() (*validator.Validate, ut.Translator) {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_trasnlations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		labelName := field.Tag.Get("label")
		return labelName
	})
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} can not be empty", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	validate.RegisterValidation("isunique", func(fl validator.FieldLevel) bool {
		params := fl.Param()
		split_params := strings.Split(params, "-")

		tableName := split_params[0]
		fieldName := split_params[1]
		fieldValue := fl.Field().String()

		return v.checkIsUnique(tableName, fieldName, fieldValue)
	})
	validate.RegisterTranslation("isunique", trans, func(ut ut.Translator) error {
		return ut.Add("isunique", "{0} already used", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isunique", fe.Field())
		return t
	})
	return validate, trans
}

func (v *Validation) Struct(s interface{}) interface{} {
	validate, trans := v.Init()
	vErrors := make(map[string]interface{})
	err := validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vErrors[e.StructField()] = e.Translate(trans)
		}
	}
	if len(vErrors) > 0 {
		return vErrors
	}
	return nil
}

func (v *Validation) checkIsUnique(tableName, fieldName, fieldValue string) bool {

	row, _ := v.conn.Query("select "+fieldName+" from "+tableName+" where "+fieldName+" =?", fieldValue)

	defer row.Close()

	var result string
	for row.Next() {
		row.Scan(&result)
	}

	return result != fieldValue
}
