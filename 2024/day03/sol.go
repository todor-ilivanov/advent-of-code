package day03

import (
	"advent/utils"
	"regexp"
)

func Solve() (int, int) {
	input := utils.ReadFileToString("2024/day03/input.txt")

	do := "do()"
	dont := "don't()"
	enabled := true

	wSize := 12

	var re = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	part1 := 0
	part2 := 0

	for i := 0; i < len(input); {
		end := utils.Min(len(input), i+wSize)
		window := input[i:end]

		if shouldEnable(window, do) {
			enabled = true
			i += len(do)
			continue
		}

		if shouldDisable(window, dont) {
			enabled = false
			i += len(dont)
			continue
		}

		matches := re.FindAllStringSubmatchIndex(window, -1)
		for _, match := range matches {

			x := utils.StringToInt(window[match[2]:match[3]])
			y := utils.StringToInt(window[match[4]:match[5]])

			part1 += x * y

			if enabled {
				part2 += x * y
			}

			i += match[1] - 1
			continue
		}

		i++
	}

	return part1, part2
}

func shouldEnable(window string, do string) bool {
	return len(window) >= len(do) && window[:len(do)] == do
}

func shouldDisable(window string, dont string) bool {
	return len(window) >= len(dont) && window[:len(dont)] == dont
}
