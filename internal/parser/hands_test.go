package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/udayangaac/poker-larvis/internal/models"
	vldtr "github.com/udayangaac/poker-larvis/internal/validator"
)

func TestHandsParser_Parse(t *testing.T) {
	parser := NewHandsParser()
	handOne := models.Hand{
		Name:  "Hand One",
		Cards: []string{"A", "A", "3", "A", "A"},
	}

	handTwo := models.Hand{
		Name:  "Hand Two",
		Cards: []string{"7", "7", "3", "7", "7"},
	}

	testValues := []struct {
		Input       string
		Hands       []models.Hand
		ExpectedErr error
	}{
		{
			Input:       "",
			Hands:       []models.Hand{},
			ExpectedErr: models.ErrInvalidInput,
		},
		{
			Input:       "Hand 1=AAQAA,Invalid Value",
			Hands:       []models.Hand{},
			ExpectedErr: models.ErrInvalidInput,
		},
		{
			Input:       "Hand 1=AAQAA,Hand 2=AAZAA",
			Hands:       []models.Hand{},
			ExpectedErr: fmt.Errorf(models.FormatErrInvalidCards, "Hand 2"),
		},
		{
			Input:       "Hand 1=AAQAA,Hand 2=AA2AAA",
			Hands:       []models.Hand{},
			ExpectedErr: fmt.Errorf(models.FormatErrInvalidNumberOfCards, "Hand 2", models.NumberOfCards),
		},
		{
			Input:       "Hand One=AA3AA,Hand Two=77377",
			Hands:       []models.Hand{handOne, handTwo},
			ExpectedErr: nil,
		},
	}

	for _, testValue := range testValues {
		hands, err := parser.Parse(testValue.Input, vldtr.NewInputStrValidator(models.NumberOfCards))
		assert.Equal(t, testValue.Hands, hands)
		assert.Equal(t, testValue.ExpectedErr, err)
	}
}
