package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordValidatorInstanse(t *testing.T) {

	t.Run("Succesful create PasswordValidator struct", func(t *testing.T) {
		expected := &PasswordValidator{
			MinLength: 5,
			MaxLength: 10,
		}
		actual := NewPasswordValidator(5, 10)

		assert.NotNil(t, actual)
		assert.Equal(t, expected, actual)
	})

	t.Run("Should return empty PasswordValidator struct when input values are not correct", func(t *testing.T) {
		expected := &PasswordValidator{}

		actual := NewPasswordValidator(-1, 10)

		assert.NotNil(t, actual)
		assert.Equal(t, expected, actual)
	})
}

func TestPasswordValidatorValidateMethod(t *testing.T) {
	validator := &PasswordValidator{MinLength: 5, MaxLength: 10}

	t.Run("Succesful password length validation", func(t *testing.T) {
		input := 8
		err := validator.Validate(&input)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})

	t.Run("Error while password length is smaller than minLength value", func(t *testing.T) {
		input := 3
		err := validator.Validate(&input)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Minimum password length - 5")
	})

	t.Run("Error while password length is greater than maxLength value", func(t *testing.T) {
		input := 11
		err := validator.Validate(&input)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Maximum password length - 10")
	})
}
