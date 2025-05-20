// Package validator contains different implementations for the validator.
package validator

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidLength       = errors.New("invalid length")
	ErrContainInvalidChars = errors.New("contain invalid characters")
)

// StringValidator creates an instance of input string validator.
func NewInputStrValidator(length int) StringValidator {
	return &inputStrValidator{
		length: length,
	}
}

type inputStrValidator struct {
	length int
}

// Validate validates the user input by checking characters and length.
func (i inputStrValidator) Validate(str string) error {
	if err := i.validateLength(str); err != nil {
		return err
	}
	return i.validateCharacters(str)
}

// validateLength validates length of the given string.
func (i inputStrValidator) validateLength(str string) error {
	arr := strings.Split(str, "")
	if len(arr) != i.length {
		return ErrInvalidLength
	}
	return nil
}

// validateCharacters validates characters of the given string.
func (i inputStrValidator) validateCharacters(str string) error {
	if !regexp.MustCompile(`^[23456789TJQKA]*$`).MatchString(str) {
		return ErrContainInvalidChars
	}
	return nil
}
