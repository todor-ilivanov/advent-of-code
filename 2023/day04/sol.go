package day04

import (
	"advent/utils"
	"math"
	"strconv"
	"strings"
)

type card struct {
	winningNumsCount int
	copies           int
}

func Solve() (int, int) {
	input := utils.ReadFileToString("2023/day04/input.txt")
	rows := strings.Split(input, "\n")

	cards := make([]card, len(rows))

	for _, row := range rows {
		cards = append(cards, createCard(row))
	}

	return calculateTotalScore(cards), calculateTotalCopies(cards)
}

func createCard(row string) card {

	cardTokens := strings.Split(row, ":")
	cardNumbers := strings.Split(cardTokens[1], "|")

	expectedNumbers := strings.TrimSpace(cardNumbers[0])
	actualNumbers := strings.TrimSpace(cardNumbers[1])

	currentCard := card{0, 1}
	expected := make(map[int]bool)

	for _, numStr := range strings.Split(expectedNumbers, " ") {
		if num, err := strconv.Atoi(numStr); err == nil {
			expected[num] = true
		}
	}

	for _, numStr := range strings.Split(actualNumbers, " ") {
		if num, err := strconv.Atoi(numStr); err == nil && expected[num] {
			currentCard.winningNumsCount += 1
		}
	}

	return currentCard
}

func calculateTotalScore(cards []card) int {

	score := 0.0

	for _, card := range cards {
		if card.winningNumsCount > 0 {
			score += math.Pow(2, float64(card.winningNumsCount-1))
		}
	}

	return int(score)
}

func calculateTotalCopies(cards []card) int {

	totalCopies := 0

	for cardIdx, card := range cards {

		totalCopies += cards[cardIdx].copies
		nextCardIdx := cardIdx + 1

		for i := nextCardIdx; i < len(cards) && i < nextCardIdx+card.winningNumsCount; i++ {
			cards[i].copies += card.copies
		}
	}

	return totalCopies
}
