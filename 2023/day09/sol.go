package day09

import (
	"advent/utils"
)

func SolvePart1(rows []string) int {

	sum := 0

	for _, row := range rows {
		initialNums := utils.ParseSpaceSeparatedInts(row)
		sum += computeLevels(initialNums, false, false)
	}

	return sum
}

func SolvePart2(rows []string) int {

	sumBackwards := 0

	for _, row := range rows {
		initialNums := utils.ParseSpaceSeparatedInts(row)
		sumBackwards += computeLevels(initialNums, false, true)
	}

	return sumBackwards
}

func computeLevels(nums []int, allZeros, backwards bool) int {

	if allZeros {
		return 0
	}

	n := len(nums)
	nextLevel := make([]int, n-1)
	allZeros = true

	for i := 0; i < n-1; i++ {
		nextLevel[i] = nums[i+1] - nums[i]
		if nextLevel[i] != 0 {
			allZeros = false
		}
	}

	if backwards {
		return nums[0] - computeLevels(nextLevel, allZeros, backwards)
	}

	return nums[n-1] + computeLevels(nextLevel, allZeros, backwards)
}
