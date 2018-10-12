package poker

// RankCategory is the type used to define hand ranking categories
type RankCategory int16

func (p RankCategory) String() string {
	return rankmap[p]
}

// Possible Poker Hand Ranking Categories
const (
	StraightFlush RankCategory = iota + 1
	FourOfAKind
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
	// To be used to represent a non value
	EmptyHand
)

var rankmap = map[RankCategory]string{
	StraightFlush: "straight-flush",
	FourOfAKind:   "four-of-a-kind",
	FullHouse:     "full-house",
	Flush:         "flush",
	Straight:      "straight",
	ThreeOfAKind:  "three-of-a-kind",
	TwoPair:       "two-pairs",
	OnePair:       "one-pair",
	HighCard:      "highest-card",
	EmptyHand:     "empty-handed :(",
}

// BetterRank return the better of two ranks
func BetterRank(rankA, rankB RankCategory) RankCategory {
	if rankA < rankB {
		return rankA
	}
	return rankB
}
