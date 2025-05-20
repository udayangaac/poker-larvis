// Package models contain all structs, types, and constants
// that have been modelled for use in poker-larvis.
package models

// Hand define the payers' hand.
type Hand struct {
	Name  string
	Cards []string
}
