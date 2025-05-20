package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/udayangaac/poker-larvis/internal/models"
)

func TestPokerDecisioner_GetDecision(t *testing.T) {
	d := NewPokerDecisioner()

	// Four of a kind
	hand0 := models.Hand{
		Name:  "Hand0",
		Cards: []string{"Q", "Q", "3", "Q", "Q"},
	}

	// Four of a kind
	hand1 := models.Hand{
		Name:  "Hand1",
		Cards: []string{"7", "7", "3", "7", "7"},
	}

	// Full house
	hand2 := models.Hand{
		Name:  "Hand2",
		Cards: []string{"K", "K", "2", "K", "2"},
	}

	// Triple
	hand3 := models.Hand{
		Name:  "Hand3",
		Cards: []string{"6", "6", "6", "3", "2"},
	}

	// Two pairs
	hand4 := models.Hand{
		Name:  "Hand4",
		Cards: []string{"7", "7", "3", "3", "2"},
	}

	// A pair
	hand5 := models.Hand{
		Name:  "Hand5",
		Cards: []string{"4", "3", "K", "9", "K"},
	}

	// High card
	hand6 := models.Hand{
		Name:  "Hand6",
		Cards: []string{"2", "9", "7", "Q", "J"},
	}

	// High card
	hand7 := models.Hand{
		Name:  "Hand7",
		Cards: []string{"2", "9", "7", "Q", "J"},
	}

	testValues := []struct {
		Hands          []models.Hand
		ExpectedWinner string
		ExpectedErr    error
	}{
		{
			Hands:          []models.Hand{hand0, hand1},
			ExpectedWinner: "Hand0",
		},
		{
			Hands:          []models.Hand{hand1, hand2},
			ExpectedWinner: "Hand1",
		},
		{
			Hands:          []models.Hand{hand2, hand3},
			ExpectedWinner: "Hand2",
		},
		{
			Hands:          []models.Hand{hand3, hand4},
			ExpectedWinner: "Hand3",
		},
		{
			Hands:          []models.Hand{hand4, hand5},
			ExpectedWinner: "Hand4",
		},
		{
			Hands:          []models.Hand{hand5, hand6},
			ExpectedWinner: "Hand5",
		},
		{
			Hands:          []models.Hand{hand6, hand7},
			ExpectedWinner: "Tie",
		},
		{
			Hands:          []models.Hand{},
			ExpectedWinner: "",
			ExpectedErr:    models.ErrNoHandsPresent,
		},
	}

	for _, testValue := range testValues {
		actualWinner, actualErr := d.GetDecision(testValue.Hands)
		assert.Equal(t, testValue.ExpectedWinner, actualWinner)
		assert.Equal(t, testValue.ExpectedErr, actualErr)
	}
}

func TestPokerDecisioner_getTopOrderHands(t *testing.T) {
	d := PokerDecisioner{cardScores: getCardScores()}

	handOne := models.Hand{
		Name:  "HandOne",
		Cards: []string{"A", "A", "A", "Q", "Q"},
	}

	handTwo := models.Hand{
		Name:  "HandTwo",
		Cards: []string{"A", "A", "A", "Q", "Q"},
	}

	handThree := models.Hand{
		Name:  "HandThree",
		Cards: []string{"A", "A", "A", "Q", "3"},
	}

	testValues := []struct {
		Hands         []models.Hand
		ExpectedHands []models.Hand
	}{
		{
			Hands:         []models.Hand{handOne, handThree},
			ExpectedHands: []models.Hand{handOne},
		},
		{
			Hands:         []models.Hand{handOne, handTwo, handThree},
			ExpectedHands: []models.Hand{handOne, handTwo},
		},
	}

	for _, testValue := range testValues {
		topHands := d.getTopOrderHands(testValue.Hands...)
		assert.EqualValues(t, testValue.ExpectedHands, topHands)
	}
}

func TestPokerDecisioner_calculateCardCombinationValue(t *testing.T) {

	d := PokerDecisioner{}

	testValues := []struct {
		Cards         []string
		ExpectedValue int
	}{
		{
			Cards:         []string{"A", "A", "A", "3", "3"},
			ExpectedValue: 13,
		},
		{
			Cards:         []string{"A", "K", "Q", "T", "9"},
			ExpectedValue: 5,
		},
	}

	for _, testValue := range testValues {
		actualValue := d.calculateCardCombinationValue(testValue.Cards)
		assert.EqualValues(t, testValue.ExpectedValue, actualValue)
	}
}

func TestPokerDecisioner_calculateScore(t *testing.T) {
	d := PokerDecisioner{cardScores: getCardScores()}

	testValues := []struct {
		Hand          models.Hand
		ExpectedScore int
	}{
		{
			Hand: models.Hand{
				Name:  "Hand",
				Cards: []string{"A", "A", "A", "Q", "Q"},
			},
			ExpectedScore: 66,
		},
	}

	for _, testValue := range testValues {
		actualScore := d.calculateScore(testValue.Hand)
		assert.Equal(t, testValue.ExpectedScore, actualScore)
	}
}
