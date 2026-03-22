package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToValueMethod(t *testing.T) {

	t.Run("Should return value from string pointer", func(t *testing.T) {
		str := "OK"
		val := ToValue(&str)

		assert.Equal(t, str, val)
	})

	t.Run("Should return value from int pointer", func(t *testing.T) {
		num := 5
		val := ToValue(&num)

		assert.Equal(t, num, val)
	})

	t.Run("Should return default type value if input pointer nil", func(t *testing.T) {
		var ptr *int
		actual := ToValue(ptr)

		assert.Equal(t, 0, actual)
	})
}

func TestIntToBoolMethod(t *testing.T) {
	t.Run("Should return True from int input while int > 0", func(t *testing.T) {
		num := 1

		val := IntToBool(&num)

		assert.True(t, val)
	})

	t.Run("Should return False from int input while int == 0 or int < 0", func(t *testing.T) {
		num := 0

		val := IntToBool(&num)

		assert.False(t, val)
	})
}
