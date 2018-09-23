package poker

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Initialization errors
var (
	ErrCharCodeInvalidLength = errors.New("The charcode is an invalid length, it must be of length 2")
	ErrInvalidCardRank       = errors.New("The provided rank is not valid, it must be either 2-9, T, J, Q, K, or A")
	ErrInvalidSuit           = errors.New("The provided suit is invalid, must be S, D, C or H")
)

var (
	rankRegEx = regexp.MustCompile("[2-9TJQKA]")
	suitRegEx = regexp.MustCompile("[SDCH]")
)

// InvalidCard is simply that - a invalid playing card. It is intended
// to be used in the event a proper card cannot be created, or another
// erroneous scenario occurs
var InvalidCard = Card{
	rankValue: 9999,
	rank:      "Uknown",
	suit:      "Uknown",
	code:      "N/A",
}

func getValue(rank string) int {
	val, err := strconv.Atoi(rank)
	if err != nil {
		switch rank {
		case "T":
			val = 10
		case "J":
			val = 11
		case "Q":
			val = 12
		case "K":
			val = 13
		case "A":
			val = 14
		}
	}
	return val
}

// Card is the primitive unit of play that makes up a hand in poker
type Card struct {
	rankValue int
	rank      string
	suit      string
	code      string
}

func (c Card) String() string {
	return fmt.Sprintf("Card: %s, Value: %d", c.code, c.rankValue)
}

// NewCard intializes a new Poker Card from a poker character cardcode.
// If initialization is not successful, the returned error value
// will report what went wrong
func NewCard(cardCode string) (Card, error) {
	if len(cardCode) != 2 {
		return InvalidCard, ErrCharCodeInvalidLength
	}

	rank := cardCode[:1]
	if !rankRegEx.MatchString(rank) {
		return InvalidCard, ErrInvalidCardRank
	}

	suit := cardCode[1:]
	if !suitRegEx.MatchString(suit) {
		return InvalidCard, ErrInvalidSuit
	}

	return Card{
		rankValue: getValue(rank),
		rank:      rank,
		suit:      suit,
		code:      cardCode,
	}, nil
}

// Hand is 5 poker cards
type Hand [5]Card

// ByRank implements sort.Interface for for a slice of Cards
type ByRank []Card

func (a ByRank) Len() int           { return len(a) }
func (a ByRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool { return a[i].rankValue < a[j].rankValue }
