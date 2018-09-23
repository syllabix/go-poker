package poker

import (
	"bytes"
	"sort"
	"strconv"
	"strings"
)

type rankcountmap map[string]int

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
	rankCounts := make(rankcountmap, numranks)

	for _, card := range hand {
		if hasFlush {
			hasFlush = flushSuit == card.suit
		}

		count, hasRank := rankCounts[card.rank]
		if hasRank {
			rankCounts[card.rank] = count + 1
		} else {
			rankCounts[card.rank] = 1
		}
	}

	return category(rankCounts, hand, hasFlush)
}

func category(rankCounts map[string]int, hand Hand, hasFlush bool) RankCategory {
	i := 0
	counts := make([]byte, len(rankCounts))
	for _, count := range rankCounts {
		counts[i] = byte(count)
		i++
	}
	sort.Slice(counts, func(i, j int) bool {
		return int(counts[i]) > int(counts[j])
	})

	if bytes.Equal(counts, quads) {
		return FourOfAKind
	} else if bytes.Equal(counts, boat) {
		return FullHouse
	} else if bytes.Equal(counts, set) {
		return ThreeOfAKind
	} else if bytes.Equal(counts, twopair) {
		return TwoPair
	} else if bytes.Equal(counts, onepair) {
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

func catKey(counts []int) string {
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	keySlice := make([]string, len(counts))
	for i, count := range counts {
		keySlice[i] = strconv.Itoa(count)
	}
	return strings.Join(keySlice, "")
}

func hasStraight(hand Hand) bool {
	sort.Sort(ByRank(hand[:]))
	return (hand[4].rankValue - hand[0].rankValue) == 4
}
