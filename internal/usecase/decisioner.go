// Package usecase implements the main logic of the poker-larvis.
package usecase

import (
	"github.com/udayangaac/poker-larvis/internal/models"
)

// NewPokerDecisioner returns an instance of Decisioner.
func NewPokerDecisioner() PokerDecisioner {
	return PokerDecisioner{
		cardScores: getCardScores(),
	}
}

type PokerDecisioner struct {
	cardScores map[string]int
}

// GetDecision decides the finial result based on the Card Combination Value
// and the Score of the card set of Hands.
func (p PokerDecisioner) GetDecision(hands []models.Hand) (string, error) {
	var topHands []models.Hand

	maxScore := 0
	topOrderHands := p.getTopOrderHands(hands...)

	for _, hand := range topOrderHands {
		s := p.calculateScore(hand)
		if s > maxScore {
			maxScore = s
			topHands = make([]models.Hand, 0)
			topHands = append(topHands, hand)
			continue
		}

		if s == maxScore {
			topHands = append(topHands, hand)
		}
	}

	if len(topHands) > 1 {
		return models.Tie, nil
	}

	if len(topHands) == 1 {
		return topHands[0].Name, nil
	}

	return "", models.ErrNoHandsPresent
}

// getTopOrderHands selects hands with top order card sets based on the Card Combination Value.
func (p PokerDecisioner) getTopOrderHands(hands ...models.Hand) []models.Hand {
	var topHands []models.Hand
	maxVal := 0
	for _, hand := range hands {
		v := p.calculateCardCombinationValue(hand.Cards)
		if maxVal < v {
			maxVal = v
			topHands = make([]models.Hand, 0)
			topHands = append(topHands, hand)
			continue
		}

		if maxVal == v {
			topHands = append(topHands, hand)
		}
	}
	return topHands
}

// calculateCardCombinationValue calculates a value for the given slice of cards.
// Calculated value can be used to identify order of card sets.
// This function calculates the sum of square of counts of each characters.
//
//	Cards: [ "A", "A", "A", "3", "2"]
//	Count of A: 3
//	Count of 3: 1
//	Count of 2: 1
//	Card Combination Value = 3^2 + 1^2 + 1^1 = 11
func (p PokerDecisioner) calculateCardCombinationValue(cards []string) int {
	val := 0
	cardCounts := make(map[string]int)

	for _, c := range cards {
		val, ok := cardCounts[c]
		if ok {
			cardCounts[c] = val + 1
			continue
		}
		cardCounts[c] = 1
	}

	for _, c := range cardCounts {
		val = val + c*c
	}
	return val
}

// calculateScore calculates the score of the given Hand.
// The score of each card are defined under getCardScores function.
func (p PokerDecisioner) calculateScore(hand models.Hand) int {
	var score int = 0
	for _, c := range hand.Cards {
		cs, ok := p.cardScores[c]
		if ok {
			score = score + cs
		}
	}
	return score
}
