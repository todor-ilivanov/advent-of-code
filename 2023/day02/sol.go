package day02

import (
	"strings"

	"advent/utils"
)

type GameSet struct {
	red   int
	green int
	blue  int
}

func Solve() (int, int) {
	input := utils.ReadFileToString("2023/day02/input.txt")
	rows := strings.Split(input, "\n")

	cubes := map[string]int{"red": 12, "green": 13, "blue": 14}
	validGamesSum := 0
	totalGamePower := 0

	for _, row := range rows {
		tokens := strings.Split(row, ":")
		gameId := parseGameId(tokens[0])
		sets := parseSets(tokens[1])

		if areAllSetsValid(cubes, sets) {
			validGamesSum += gameId
		}

		totalGamePower += calculateGamePower(sets)
	}

	return validGamesSum, totalGamePower
}

func parseGameId(game string) int {
	gameId := strings.Split(game, " ")[1]
	return utils.StringToInt(gameId)
}

func parseSets(setsStr string) []GameSet {
	sets := strings.Split(setsStr, ";")

	var gameSets []GameSet

	for _, set := range sets {
		trimmedSet := strings.TrimSpace(set)
		picks := strings.Split(trimmedSet, ",")

		var gameSet GameSet

		for _, pick := range picks {
			trimmedPick := strings.TrimSpace(pick)
			pickTokens := strings.Split(trimmedPick, " ")

			if pickTokens[1] == "red" {
				gameSet.red += utils.StringToInt(pickTokens[0])
			}
			if pickTokens[1] == "green" {
				gameSet.green += utils.StringToInt(pickTokens[0])
			}
			if pickTokens[1] == "blue" {
				gameSet.blue += utils.StringToInt(pickTokens[0])
			}
		}

		gameSets = append(gameSets, gameSet)
	}
	return gameSets
}

func areAllSetsValid(cubes map[string]int, sets []GameSet) bool {
	for _, set := range sets {
		if set.red > cubes["red"] || set.green > cubes["green"] || set.blue > cubes["blue"] {
			return false
		}
	}
	return true
}

func calculateGamePower(sets []GameSet) int {
	maxRed, maxGreen, maxBlue := 0, 0, 0

	for _, set := range sets {
		maxRed = utils.Max(maxRed, set.red)
		maxGreen = utils.Max(maxGreen, set.green)
		maxBlue = utils.Max(maxBlue, set.blue)
	}

	return maxRed * maxGreen * maxBlue
}
