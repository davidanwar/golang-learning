package golang_validation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		panic("validate fail")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	name := ""
	err := validate.Var(name, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariable(t *testing.T) {
	validate := validator.New()
	password := "password"
	confirmPassword := "passwordsalah"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "david"
	err := validate.Var(user, "required,numeric")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	user := "100"
	err := validate.Var(user, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationStruct(t *testing.T) {
	type User struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	user := User{
		Username: "david",
		Password: "password",
	}
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	type User struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	user := User{
		Username: "david",
		Password: "password",
	}
	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validation := range validationErrors {
			fmt.Println("Error ", validation.Field(), "On Tag ", validation.Tag(), "with Error ", validation.Error())
		}
		fmt.Println(err.Error())
	}
}
