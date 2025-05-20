// Package usecase implements the main logic of the poker-larvis.
package usecase

// getCardScores returns a map which contains card as the Key
// and value of the card as the Value.
func getCardScores() map[string]int {
	return map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
}
