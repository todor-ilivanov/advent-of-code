package day05

import (
	"regexp"
	"strings"

	"advent/day05/dayutils"
	"advent/utils"
)

func SolvePart1() string {
	stacks := dayutils.InitStacksFromFile("day05/init-stacks.txt")
	movesStr := utils.ReadFileToString("day05/input.txt")
	moves := strings.Split(movesStr, "\n")

	for _, move := range moves {
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(move, -1)
		moveCount := utils.StringToInt(nums[0])
		from := utils.StringToInt(nums[1]) - 1
		to := utils.StringToInt(nums[2]) - 1

		for i := 0; i < moveCount; i++ {
			topEl := peek(stacks[from])
			stacks[from] = pop(stacks[from])
			stacks[to] = append(stacks[to], topEl)
		}
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

func peek(stack []string) string {
	if len(stack) == 0 {
		return ""
	}
	if len(stack) == 1 {
		return stack[0]
	}
	return stack[len(stack)-1]
}

func pop(stack []string) []string {
	if len(stack) == 0 {
		return make([]string, 0)
	}

	return stack[:len(stack)-1]
}
