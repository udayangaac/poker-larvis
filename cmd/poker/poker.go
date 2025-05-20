package main

import (
	"fmt"

	"os"

	"github.com/udayangaac/poker-larvis/internal/models"
	"github.com/udayangaac/poker-larvis/internal/parser"
	"github.com/udayangaac/poker-larvis/internal/usecase"
	vldtr "github.com/udayangaac/poker-larvis/internal/validator"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	handsStr = kingpin.
		Flag("hands",
			"Define hands in the game. Example: poker --hands=smith=QQA78,'scott nelson'=87A22. "+
				"Note: String with spaces should be wrapped with apostrophes.").
		Short('h').
		PlaceHolder("NAME_ONE=CARD_SET_ONE,NAME_TWO=CARD_SET_TWO,...").
		Required().String()
)

func main() {

	kingpin.Version("0.0.1")
	kingpin.Parse()

	// Create a validator to validate the input.
	inputStrValidator := vldtr.NewInputStrValidator(models.NumberOfCards)

	// Create a parser to validates and extracts hands from the input.
	parser := parser.NewHandsParser()
	hands, err := parser.Parse(*handsStr, inputStrValidator)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	// Creates a decisioner to take the decision of given hands.
	decisioner := usecase.NewPokerDecisioner()
	decision, err := decisioner.GetDecision(hands)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Winner: %s\n", decision)
}
