// Package models contain all structs, types, and constants
// that have been modelled for use in poker-larvis.
package models

import "errors"

var (
	// ErrNoHandsPresent If no cards were found, this value is returned.
	ErrNoHandsPresent = errors.New("no hands present")

	// ErrInvalidInput If the input was invalid, this error is returned.
	ErrInvalidInput = errors.New("invalid input")

	// FormatErrInvalidNumberOfCards is used to generate an error if the number of cards does not match.
	FormatErrInvalidNumberOfCards string = "number of cards in %v should be %v"

	// FormatErrInvalidCards is used to generate an error if an invalid string was used to define cards.
	FormatErrInvalidCards string = "invalid character! use only 23456789TJQKA to define cards in %v"
)
