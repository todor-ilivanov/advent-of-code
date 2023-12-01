package day02

import (
	"fmt"
	"strings"

	"advent/day02/dayutils"
	"advent/utils"
)

var moves = [3]string{"r", "p", "s"}

const winScore int = 6
const drawScore int = 3

func SolvePart1() int {

	str := utils.ReadFileToString("day02/input.txt")

	var decryptMap = make(map[string][2]string)

	decryptMap[moves[0]] = [2]string{"A", "X"}
	decryptMap[moves[1]] = [2]string{"B", "Y"}
	decryptMap[moves[2]] = [2]string{"C", "Z"}

	var sum int

	rounds := strings.Split(str, "\n")
	for _, round := range rounds {
		moves := strings.Split(round, " ")
		them := decryptMove(moves[0], decryptMap)
		us := decryptMove(moves[1], decryptMap)
		score := calcScoreForRound(them, us)
		sum += score
	}
	return sum
}

func calcScoreForRound(them, us string) int {

	myScore := dayutils.GetScoreForMove(us)

	if them == us {
		return myScore + drawScore
	}

	if isWin(them, us) {
		return myScore + winScore
	}

	return myScore
}

func isWin(theirMove, ourMove string) bool {
	return (ourMove == "r" && theirMove == "s") || (ourMove == "p" && theirMove == "r") || (ourMove == "s" && theirMove == "p")
}

func decryptMove(encrypted string, decryptMap map[string][2]string) string {
	for _, move := range moves {
		if isElementIn(encrypted, decryptMap[move]) {
			return move
		}
	}
	fmt.Println("Unrecognized move.")
	return "unknown"
}

func isElementIn(move string, moveArr [2]string) bool {
	for _, el := range moveArr {
		if el == move {
			return true
		}
	}
	return false
}
