package player

import (
	"fmt"
	"strings"

	"github.com/syllabix/psychic-poker-player/poker"
)

func assembleHand(codes []string) (poker.Hand, error) {
	hand := poker.Hand{}
	for i, code := range codes {
		card, err := poker.NewCard(code)
		if err != nil {
			return hand, err
		}
		hand[i] = *card
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

func RevealBestHand(input string) (string, error) {
	cardCodes := strings.Split(input, " ")

	hand := cardCodes[:5]
	deck := cardCodes[5:]

	bestHand, err := findBestHand(hand, deck)
	if err != nil {
		return "An error occurred while examining cards", err
	}
	return fmt.Sprintf("Hand: %s Deck: %s Best hand: %s", strings.Join(hand, " "), strings.Join(deck, " "), bestHand), nil
}
