package day06

import (
	"advent/utils"
	"strings"
)

func SolvePart1() int {
	input := utils.ReadFileToString("2023/day06/input.txt")
	rows := strings.Split(input, "\n")

	durations := convertRowToIntSlice(rows[0])
	targets := convertRowToIntSlice(rows[1])
	numRaces := len(durations)

	total := 1

	for i := 0; i < numRaces; i++ {
		duration := durations[i]
		target := targets[i]
		bestL, bestR := binarySearchForMinAndMax(duration, target)
		total *= bestR - bestL - 1
	}

	return total
}

func SolvePart2() int {

	input := utils.ReadFileToString("2023/day06/input.txt")
	rows := strings.Split(input, "\n")

	durationStr := strings.Split(rows[0], " ")
	targetStr := strings.Split(rows[1], " ")

	// faster to hardcode them, only doing parsing to obstruct the input data
	duration := utils.StringToInt(strings.Join(durationStr, ""))
	target := utils.StringToInt(strings.Join(targetStr, ""))

	bestL, bestR := binarySearchForMinAndMax(duration, target)

	return bestR - bestL - 1
}

func binarySearchForMinAndMax(duration, target int) (int, int) {
	var min, max int

	start, end := 0, duration

	for start <= end {
		mid := start + (end-start)/2
		currentDuration := mid * (duration - mid)

		if currentDuration <= target {
			min = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	start, end = 0, duration

	for start <= end {
		mid := start + (end-start)/2
		currentDuration := mid * (duration - mid)

		if currentDuration <= target {
			max = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return min, max
}

func convertRowToIntSlice(row string) []int {
	slice := make([]int, 0)

	for _, numStr := range strings.Split(row, " ") {
		num := utils.StringToInt(numStr)
		slice = append(slice, num)
	}

	return slice
}
