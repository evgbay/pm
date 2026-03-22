package util

import "fmt"

type PasswordValidator struct {
	MinLength int
	MaxLength int
}

func (pv *PasswordValidator) Validate(passwordLength *int) error {
	len := ToValue(passwordLength)
	if len < pv.MinLength {
		return fmt.Errorf("Minimum password length - %d", pv.MinLength)
	}
	if len > pv.MaxLength {
		return fmt.Errorf("Maximum password length - %d", pv.MaxLength)
	}
	return nil
}

func NewPasswordValidator(minLength int, maxLength int) *PasswordValidator {
	if minLength > Zero && maxLength <= MaxPasswordLength {
		return &PasswordValidator{
			MinLength: minLength,
			MaxLength: maxLength,
		}
	}
	return &PasswordValidator{}
}
