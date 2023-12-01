package day01

import (
	"strings"
	"unicode"

	"advent/utils"
)

func Solve() (int, int) {
	input := utils.ReadFileToString("2023/day01/input.txt")
	rows := strings.Split(input, "\n")

	sumPart1 := 0
	sumPart2 := 0

	for _, row := range rows {
		sumPart1 += GetDigitPart1(row)
		sumPart2 += GetDigitPart2(row)
	}

	return sumPart1, sumPart2
}

func GetDigitPart1(row string) int {

	lDigit, rDigit := 0, 0
	l, r := 0, len(row)-1

	for l <= r && (lDigit == 0 || rDigit == 0) {

		lRune, rRune := rune(row[l]), rune(row[r])

		if !unicode.IsDigit(lRune) {
			l += 1
		} else if lDigit == 0 {
			lDigit = int(lRune) - '0'
		}

		if !unicode.IsDigit(rRune) {
			r -= 1
		} else if rDigit == 0 {
			rDigit = int(rRune) - '0'
		}
	}

	return 10*lDigit + rDigit
}

func GetDigitPart2(row string) int {

	digitStrings := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	lDigit, rDigit := 0, 0
	l, r := 0, len(row)-1

	for l <= r && (lDigit == 0 || rDigit == 0) {

		for i := 3; i <= 5; i++ {

			if lDigit == 0 && l+i <= len(row) {
				windowL := row[l : l+i]
				if digit, ok := digitStrings[windowL]; ok {
					lDigit = digit
				}
			}

			if rDigit == 0 && r-i+1 >= 0 {
				windowR := row[r-i+1 : r+1]
				if digit, ok := digitStrings[windowR]; ok {
					rDigit = digit
				}
			}
		}

		lRune, rRune := rune(row[l]), rune(row[r])

		if !unicode.IsDigit(lRune) {
			l += 1
		} else if lDigit == 0 {
			lDigit = int(lRune) - '0'
		}

		if !unicode.IsDigit(rRune) {
			r -= 1
		} else if rDigit == 0 {
			rDigit = int(rRune) - '0'
		}
	}

	return 10*lDigit + rDigit
}
