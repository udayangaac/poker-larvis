// Package validator contains different implementations for the validator.
package validator

// StringValidator is a generic interface for string validations.
type StringValidator interface {

	// Validate validates the given string and if it was invalid,
	// it will return an error according to the implementation.
	Validate(str string) error
}
