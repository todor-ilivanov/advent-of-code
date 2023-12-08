package day07

import (
	"advent/utils"
	"sort"
	"strings"
)

var CARDS = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type hand struct {
	rank    int
	handStr string
	bid     int
}

func newHand(handStr string, bid int) *hand {
	charCounts := make(map[rune]int)

	for _, char := range handStr {
		charCounts[char]++
	}

	pairs, triplets := 0, 0

	for _, v := range charCounts {
		switch v {
		case 5:
			return &hand{6, handStr, bid}
		case 4:
			return &hand{5, handStr, bid}
		case 3:
			triplets++
		case 2:
			pairs++
		}
	}

	rank := determineCardRank(triplets, pairs)
	return &hand{rank, handStr, bid}
}

func newHandWithJoker(handStr string, bid int) *hand {
	charCounts := make(map[rune]int)
	jCount := 0
	maxFreq := 0

	for _, char := range handStr {
		if char == 'J' {
			jCount++
		} else {
			charCounts[char]++
		}
		maxFreq = utils.Max(maxFreq, charCounts[char])
	}

	pairs, triplets := 0, 0
	for _, v := range charCounts {
		switch v {
		case 3:
			triplets++
		case 2:
			pairs++
		}
	}

	switch maxFreq + jCount {
	case 5:
		return &hand{6, handStr, bid}
	case 4:
		return &hand{5, handStr, bid}
	case 3:
		if jCount > 0 {
			triplets++
			pairs = utils.Max(0, pairs-1)
		}
	case 2:
		if jCount > 0 {
			pairs++
		}
	}

	rank := determineCardRank(triplets, pairs)
	return &hand{rank, handStr, bid}
}

func Solve() (int, int) {
	input := utils.ReadFileToString("2023/day07/input.txt")
	rows := strings.Split(input, "\n")

	var hands []*hand
	var handsWithJoker []*hand

	for _, row := range rows {
		tokens := strings.Split(row, " ")
		handStr := tokens[0]
		bid := utils.StringToInt(tokens[1])

		hand := newHand(handStr, bid)
		hands = append(hands, hand)

		handWithJoker := newHandWithJoker(handStr, bid)
		handsWithJoker = append(handsWithJoker, handWithJoker)
	}

	sortHands(hands, false)

	total := 0
	for i := 0; i < len(hands); i++ {
		total += (i + 1) * hands[i].bid
	}

	sortHands(handsWithJoker, true)

	totalWithJoker := 0
	for i := 0; i < len(handsWithJoker); i++ {
		totalWithJoker += (i + 1) * handsWithJoker[i].bid
	}

	return total, totalWithJoker
}

func determineCardRank(triplets, pairs int) int {
	if pairs == 1 && triplets == 1 {
		return 4
	} else if triplets == 1 {
		return 3
	}
	return pairs
}

func sortHands(hands []*hand, joker bool) {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].rank == hands[j].rank {
			return compareHands(hands[i].handStr, hands[j].handStr, joker)
		}
		return hands[i].rank < hands[j].rank
	})
}

func compareHands(h1, h2 string, joker bool) bool {
	for i := 0; i < len(h1); i++ {
		if h1[i] == h2[i] {
			continue
		}
		r1, r2 := rune(h1[i]), rune(h2[i])
		return getCardPower(r1, joker) < getCardPower(r2, joker)
	}
	return true
}

func getCardPower(card rune, joker bool) int {
	if card == 'J' && joker {
		return 1
	}
	return CARDS[card]
}
