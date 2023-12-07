package day07

import (
	"advent/utils"
	"fmt"
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

	rank, pairs, triplets := 0, 0, 0

	for _, v := range charCounts {
		switch v {
		case 5:
			rank = 6
		case 4:
			rank = 5
		case 3:
			triplets += 1
		case 2:
			pairs += 1
		}

		if rank > 0 {
			break
		}
	}

	if pairs == 1 && triplets == 1 {
		rank = 4
	} else if triplets == 1 {
		rank = 3
	} else if rank == 0 {
		rank = pairs
	}

	return &hand{rank, handStr, bid}
}

func newHand2(handStr string, bid int) *hand {
	return &hand{}
}

func Solve() (int, int) {
	input := utils.ReadFileToString("2023/day07/input.txt")
	rows := strings.Split(input, "\n")

	var hands []*hand

	for _, row := range rows {
		tokens := strings.Split(row, " ")
		handStr := tokens[0]
		bid := utils.StringToInt(tokens[1])

		hand := newHand(handStr, bid)
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].rank == hands[j].rank {
			return compareHands(hands[i].handStr, hands[j].handStr)
		}
		return hands[i].rank < hands[j].rank
	})

	total := 0
	for i := 0; i < len(hands); i++ {
		total += (i + 1) * hands[i].bid
		fmt.Println(hands[i].rank, hands[i].handStr, hands[i].bid)
	}

	return total, 0
}

func compareHands(h1, h2 string) bool {
	for i := 0; i < len(h1); i++ {
		if h1[i] == h2[i] {
			continue
		}
		r1, r2 := rune(h1[i]), rune(h2[i])
		return CARDS[r1] < CARDS[r2]
	}
	return true
}
