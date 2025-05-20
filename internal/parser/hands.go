// Package parser is a collection of parsers that are used to parse user inputs.
package parser

import (
	"fmt"
	"strings"

	"github.com/udayangaac/poker-larvis/internal/models"
	vldtr "github.com/udayangaac/poker-larvis/internal/validator"
)

const (
	separator   = ","
	kvSeparator = "="
)

// NewHandsParser creates an instance of HandsParser.
func NewHandsParser() HandsParser {
	return HandsParser{}
}

type HandsParser struct{}

// Parse extracts hands from the input and checks their length and characters.
func (h HandsParser) Parse(input string, validators ...vldtr.StringValidator) ([]models.Hand, error) {
	hands := make([]models.Hand, 0)
	pairs := strings.Split(input, separator)

	if input == "" {
		return []models.Hand{}, models.ErrInvalidInput
	}

	for _, pair := range pairs {

		kv := strings.Split(pair, kvSeparator)

		if len(kv) != 2 {
			return []models.Hand{}, models.ErrInvalidInput
		}

		name := kv[0]
		cardsStr := kv[1]

		for _, validator := range validators {
			if err := validator.Validate(cardsStr); err != nil {
				if err == vldtr.ErrInvalidLength {
					err = fmt.Errorf(models.FormatErrInvalidNumberOfCards, name, models.NumberOfCards)
				}
				if err == vldtr.ErrContainInvalidChars {
					err = fmt.Errorf(models.FormatErrInvalidCards, name)
				}
				return []models.Hand{}, err
			}
		}

		cards := strings.Split(cardsStr, "")
		hands = append(hands, models.Hand{
			Name:  name,
			Cards: cards,
		})
	}

	return hands, nil
}
