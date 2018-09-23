package poker

import (
	"errors"
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

// Card is the primative unit of play that makes up a hand in poker
type Card struct {
	rankValue int
	rank      string
	suit      string
	code      string
}

func (c Card) String() string {
	return c.code
}

// NewCard intializes a new Poker Card from a poker character cardcode.
// If initialization is not successful, the returned error value
// will report what went wrong
func NewCard(cardCode string) (*Card, error) {
	if len(cardCode) != 2 {
		return nil, ErrCharCodeInvalidLength
	}

	rank := cardCode[:1]
	if !rankRegEx.MatchString(rank) {
		return nil, ErrInvalidCardRank
	}

	suit := cardCode[1:]
	if !suitRegEx.MatchString(suit) {
		return nil, ErrInvalidSuit
	}

	return &Card{
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
