package poker

import (
	"bytes"
	"math"
	"sort"
)

const (
	numranks = 13
)

var (
	quads   = []byte{byte(4), byte(1)}
	boat    = []byte{byte(3), byte(2)}
	set     = []byte{byte(3), byte(1), byte(1)}
	twopair = []byte{byte(2), byte(2), byte(1)}
	onepair = []byte{byte(2), byte(1), byte(1), byte(1)}
)

// GetRank takes a poker hand and returns its HandRank
func GetRank(hand Hand) RankCategory {
	hasFlush := true
	flushSuit := hand[0].suit
	rankCounts := makeRankCounter()

	for _, card := range hand {
		if hasFlush {
			hasFlush = flushSuit == card.suit
		}

		rankCounts[card.rankValue-2]++
	}

	return category(rankCounts, hand, hasFlush)
}

func category(counts []byte, hand Hand, hasFlush bool) RankCategory {

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	if bytes.HasPrefix(counts, quads) {
		return FourOfAKind
	} else if bytes.HasPrefix(counts, boat) {
		return FullHouse
	} else if bytes.HasPrefix(counts, set) {
		return ThreeOfAKind
	} else if bytes.HasPrefix(counts, twopair) {
		return TwoPair
	} else if bytes.HasPrefix(counts, onepair) {
		return OnePair
	}

	if hasStraight(hand) {
		if hasFlush {
			return StraightFlush
		}
		return Straight
	}

	if hasFlush {
		return Flush
	}

	return HighCard
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func hasStraight(hand Hand) bool {
	maxValue := math.MinInt16
	minValue := math.MaxInt16
	for _, card := range hand {
		maxValue = max(card.rankValue, maxValue)
		minValue = min(card.rankValue, minValue)
	}
	return (maxValue - minValue) == 4
}

func makeRankCounter() []byte {
	counter := make([]byte, 13, 13)
	for idx := range counter {
		counter[idx] = 0
	}
	return counter
}
