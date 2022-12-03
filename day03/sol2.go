package day03

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/hashicorp/go-set"
)

const groupSize int = 3

func SolvePart2() int {
	b, err := os.ReadFile("day03/input.txt")

	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	rucksacks := strings.Split(str, "\n")

	var sum int

	for i := 0; i < len(rucksacks); i += groupSize {
		group := rucksacks[i : i+groupSize]
		var countArray [52]int

		for _, rucksack := range group {
			seen := set.New[rune](52)
			for _, ch := range rucksack {
				idx := getIdxForChar(ch)
				if !seen.Contains(ch) {
					countArray[idx] += 1
					seen.Insert(ch)
				}
			}
		}

		for i, count := range countArray {
			if count == groupSize {
				sum += i + 1
				break
			}
		}
	}

	return sum
}

func getIdxForChar(ch rune) int {
	if unicode.IsUpper(ch) {
		return getIdxForUpperCase(ch)
	} else if unicode.IsLower(ch) {
		return getIdxForLowerCase(ch)
	}

	return 0
}

func getIdxForLowerCase(charRune rune) int {
	return int(charRune) - 97
}

func getIdxForUpperCase(charRune rune) int {
	return int(charRune) - 39
}
