package poker

import (
	"sort"
	"strconv"
	"strings"
)

type rankcountmap map[string]int

const (
	numranks = 13
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
	counts := make([]int, len(rankCounts))
	for _, count := range rankCounts {
		counts[i] = count
		i++
	}
	key := catKey(counts)

	switch key {
	case "41":
		return FourOfAKind
	case "32":
		return FullHouse
	case "311":
		return ThreeOfAKind
	case "221":
		return TwoPair
	case "2111":
		return OnePair
	default:
		break
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
