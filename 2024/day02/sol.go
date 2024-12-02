package day02

import (
	"advent/utils"
	"strings"
)

func Solve() (int, int) {
	input := utils.ReadFileToString("2024/day02/input.txt")
	rows := strings.Split(input, "\n")

	var safeReportsPart1 [][]int
	var safeReportsPart2 [][]int

	for _, row := range rows {
		report := utils.ParseSpaceSeparatedInts(row)

		idxToRemove := isSafe(report)

		if idxToRemove == -1 {
			safeReportsPart1 = append(safeReportsPart1, report)
		} else {

			for i := range report {
				if isSafe(removeElement(report, i)) == -1 {
					safeReportsPart2 = append(safeReportsPart2, report)
					break
				}
			}
		}
	}

	return len(safeReportsPart1), len(safeReportsPart1) + len(safeReportsPart2)
}

func isSafe(report []int) int {
	prev := report[0]

	isIncreasing := report[1]-prev > 0

	for i := 1; i < len(report); i++ {

		diff := report[i] - prev

		if isDirMismatch(isIncreasing, diff) || isOutsideRange(diff) {
			return i
		}

		prev = report[i]
	}

	return -1
}

func removeElement(slice []int, index int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice
}

func isDirMismatch(isIncreasing bool, diff int) bool {
	return isIncreasing == (diff < 0)
}

func isOutsideRange(num int) bool {
	numAbs := utils.Abs(num)
	return numAbs > 3 || numAbs < 1
}
