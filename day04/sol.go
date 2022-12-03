package day04

import (
	"fmt"
	"os"
	"strings"

	"advent/utils"
)

func Solve() (int, int) {
	b, err := os.ReadFile("day04/input.txt")

	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	rows := strings.Split(str, "\n")

	var part1 int
	var part2 int

	for _, row := range rows {
		assignments := strings.Split(row, ",")
		elf1 := strings.Split(assignments[0], "-")
		elf2 := strings.Split(assignments[1], "-")

		elf1Lower := utils.StringToInt(elf1[0])
		elf1Upper := utils.StringToInt(elf1[1])
		elf2Lower := utils.StringToInt(elf2[0])
		elf2Upper := utils.StringToInt(elf2[1])

		if overlapsFully(elf1Lower, elf1Upper, elf2Lower, elf2Upper) {
			part1 += 1
		}

		if overlaps(elf1Lower, elf1Upper, elf2Lower, elf2Upper) {
			part2 += 1
		}
	}
	return part1, part2
}

func overlapsFully(lower, upper, lowerOther, upperOther int) bool {
	return (lower >= lowerOther && upper <= upperOther) || (lowerOther >= lower && upperOther <= upper)
}

func overlaps(lower, upper, lowerOther, upperOther int) bool {
	return (lower >= lowerOther && lower <= upperOther) || (lowerOther >= lower && lowerOther <= upper)
}
