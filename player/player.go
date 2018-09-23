package player

import (
	"errors"
	"fmt"
	"strings"

	"github.com/syllabix/psychic-poker-player/poker"
)

// Errors
var (
	ErrInvalidInput = errors.New("The provided input is invalid. Expecting a total of 10 card codes delimited by a space")
)

func assembleHand(codes []string) (poker.Hand, error) {
	hand := poker.Hand{}
	for i, code := range codes {
		card, err := poker.NewCard(code)
		if err != nil {
			return hand, err
		}
		hand[i] = card
	}
	return hand, nil
}

func findBestHand(handCodes, deckCodes []string) (poker.RankCategory, error) {
	hand, err := assembleHand(handCodes)
	if err != nil {
		return poker.EmptyHand, err
	}
	deck, err := assembleHand(deckCodes)
	if err != nil {
		return poker.EmptyHand, err
	}
	bestRank := poker.BetterRank(
		poker.GetRank(hand),
		poker.GetRank(deck))

	newHand := make([]string, 5, 5)
	for i := 0; i < 4; i++ {
		drawnCards := deckCodes[:i+1]

		for hindex := range handCodes {
			copy(newHand, handCodes)
			for dindex, drawnCard := range drawnCards {
				idx := hindex + dindex
				if idx > 4 {
					idx = idx - 5
				}
				newHand[idx] = drawnCard
			}
			h, err := assembleHand(newHand)
			if err != nil {
				return poker.EmptyHand, err
			}
			curRank := poker.GetRank(h)
			bestRank = poker.BetterRank(curRank, bestRank)
		}
	}

	return bestRank, nil
}

// RevealBestHand takes a string of input representing a 5 card poker hand
// as well as the top 5 cards in a deck.
// The string itself should be made up of space delimited
// card codes. The code format is comprized of 2 characters symbolizing Rank and Suit:
//
// Ranks:
// 2-9
// T = 10,
// J = Jack
// Q = Queen
// K = King
// A = Ace
//
// Suits:
// S = Spades
// C = Clubs
// D = Diamonds
// H = Hearts
//
// Example Input:
// 2H 2S 3H 3S 3C 2D 9C 3D 6C TH
// ^------------^ ^------------^
//      Hand           Deck
//
// Result:
// Hand: 2H 2S 3H 3S 3C Deck: 2D 9C 3D 6C TH Best hand: full-house
func RevealBestHand(input string) (string, error) {
	cardCodes := strings.Split(input, " ")

	if len(cardCodes) != 10 {
		return "", ErrInvalidInput
	}

	hand := cardCodes[:5]
	deck := cardCodes[5:]

	bestHand, err := findBestHand(hand, deck)
	if err != nil {
		return "An error occurred while examining cards", err
	}

	return fmt.Sprintf("Hand: %s Deck: %s Best hand: %s",
		strings.Join(hand, " "),
		strings.Join(deck, " "),
		bestHand), nil
}
