package day03

import (
	"strings"
	"unicode"

	"advent/utils"

	"github.com/hashicorp/go-set"
)

func SolvePart1() int {

	str := utils.ReadFileToString("2022/day03/input.txt")
	rucksacks := strings.Split(str, "\n")

	var sum int

	for _, rucksack := range rucksacks {

		midIdx := len(rucksack) / 2

		seenChars1 := set.From(strings.Split(rucksack[0:midIdx], ""))
		seenChars2 := set.From(strings.Split(rucksack[midIdx:], ""))

		for _, ch := range rucksack {
			charStr := string(ch)

			if seenChars1.Contains(charStr) && seenChars2.Contains(charStr) {
				sum += calcPrioForChar(ch)
			}
			seenChars1.Remove(charStr)
			seenChars2.Remove(charStr)
		}
	}
	return sum
}

func calcPrioForChar(ch rune) int {
	if unicode.IsUpper(ch) {
		return upperCaseToPrio(ch)
	} else if unicode.IsLower(ch) {
		return lowerCaseToPrio(ch)
	}

	return 0
}

func lowerCaseToPrio(charRune rune) int {
	return int(charRune) - 96
}

func upperCaseToPrio(charRune rune) int {
	return int(charRune) - 38
}
