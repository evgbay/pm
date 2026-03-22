package model

import (
	"github.com/evgbay/pm/util"
	"log"
)

type Template struct {
	Len     int
	Letters bool
	Digits  bool
	Special bool
}

func NewTemplate(passwordLength *int, enableSpecialSymbols *int) *Template {
	validator := util.NewPasswordValidator(util.MinPasswordLength, util.MaxPasswordLength)
	err := validator.Validate(passwordLength)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Template{
		Len:     util.ToValue(passwordLength),
		Letters: true,
		Digits:  true,
		Special: util.IntToBool(enableSpecialSymbols),
	}
}
