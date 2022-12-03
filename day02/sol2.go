package day02

import (
	"strings"

	"advent/day02/dayutils"
	"advent/utils"
)

const pointsForWin int = 6
const pointsForDraw int = 3

func SolvePart2() int {
	str := utils.ReadFileToString("day02/input.txt")

	var decryptMap = make(map[string]string)

	decryptMap["A"] = "r"
	decryptMap["B"] = "p"
	decryptMap["C"] = "s"

	var sum int

	rounds := strings.Split(str, "\n")
	for _, round := range rounds {
		moves := strings.Split(round, " ")
		them := decryptMap[moves[0]]
		result := moves[1]
		score := calcScoreForExpectedResult(them, result)
		sum += score
	}
	return sum
}

func calcScoreForExpectedResult(them, result string) int {

	if result == "X" {
		return getLoseMovePoints(them)
	}

	if result == "Y" {
		return dayutils.GetScoreForMove(them) + pointsForDraw
	}

	if result == "Z" {
		return getWinMovePoints(them) + pointsForWin
	}

	return 0
}

func getLoseMovePoints(theirMove string) int {
	if theirMove == "r" {
		return dayutils.GetScoreForMove("s")
	} else if theirMove == "p" {
		return dayutils.GetScoreForMove("r")
	} else {
		return dayutils.GetScoreForMove("p")
	}
}

func getWinMovePoints(theirMove string) int {
	if theirMove == "r" {
		return dayutils.GetScoreForMove("p")
	} else if theirMove == "p" {
		return dayutils.GetScoreForMove("s")
	} else {
		return dayutils.GetScoreForMove("r")
	}
}
