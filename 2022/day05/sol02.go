package day05

import (
	"regexp"
	"strings"

	"advent/day05/dayutils"
	"advent/utils"
)

func SolvePart2() string {
	stacks := dayutils.InitStacksFromFile("day05/init-stacks.txt")
	movesStr := utils.ReadFileToString("day05/input.txt")
	moves := strings.Split(movesStr, "\n")

	for _, move := range moves {
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(move, -1)
		moveCount := utils.StringToInt(nums[0])
		from := utils.StringToInt(nums[1]) - 1
		to := utils.StringToInt(nums[2]) - 1

		cutoff := len(stacks[from]) - moveCount
		crates := stacks[from][cutoff:]
		stacks[from] = stacks[from][:cutoff]
		stacks[to] = append(stacks[to], crates...)
	}

	var result []string

	for _, stack := range stacks {
		if len(stack) > 0 {
			lastIdx := len(stack) - 1
			result = append(result, stack[lastIdx])
		}
	}

	return strings.Join(result, "")
}
