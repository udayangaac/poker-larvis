package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputValidator_Validate(t *testing.T) {
	v := NewInputStrValidator(5)

	testValues := []struct {
		Str         string
		ExpectedErr error
	}{
		{
			Str:         "22AAAA", // Contains six characters.
			ExpectedErr: ErrInvalidLength,
		},
		{
			Str:         "2233U", // Invalid character present in the string.
			ExpectedErr: ErrContainInvalidChars,
		},
		{
			Str:         "3233T",
			ExpectedErr: nil,
		},
	}

	for _, testValue := range testValues {
		actualError := v.Validate(testValue.Str)
		assert.Equalf(t, testValue.ExpectedErr, actualError, "Input string: %v", testValue.Str)
	}
}
