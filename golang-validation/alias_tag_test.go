package golang_validation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestAliasTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,min=5,max=100")

	type Seller struct {
		Name    string `validate:"varchar"`
		Address string `validate:"varchar"`
	}

	seller := Seller{
		Name:    "seller",
		Address: "addhgf",
	}

	err := validate.Struct(seller)
	if err != nil {
		fmt.Println(err.Error())
	}
}
