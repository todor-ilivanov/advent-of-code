package day06

import (
	"advent/utils"
)

func SolveForWindowSize(windowSize int) int {

	readings := utils.ReadFileToString("day06/input.txt")
	charCounts := make(map[rune]int)

	for i := 0; i < windowSize; i++ {
		char := []rune(readings)[i]
		charCounts = incrementKey(char, charCounts)
	}

	var result int = windowSize

	for i := windowSize; i < len(readings); i++ {
		if len(charCounts) == windowSize {
			break
		}
		curr := []rune(readings)[i]
		charCounts = incrementKey(curr, charCounts)
		prev := []rune(readings)[i-windowSize]
		charCounts[prev] -= 1

		if charCounts[prev] == 0 {
			delete(charCounts, prev)
		}
		result++
	}

	return result
}

func incrementKey(key rune, charMap map[rune]int) map[rune]int {
	if _, isPresent := charMap[key]; isPresent {
		charMap[key] += 1
	} else {
		charMap[key] = 1
	}
	return charMap
}
