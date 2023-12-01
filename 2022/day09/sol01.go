package day09

import (
	"strings"

	"advent/utils"
)

type pos struct {
	row int
	col int
}

func SolvePart1() int {
	input := utils.ReadFileToString("2022/day09/input.txt")
	rows := strings.Split(input, "\n")

	dirMap := make(map[string]pos)
	dirMap["R"] = pos{0, 1}
	dirMap["L"] = pos{0, -1}
	dirMap["U"] = pos{-1, 0}
	dirMap["D"] = pos{1, 0}

	visitedPos := make(map[pos]int)
	head := pos{0, 0}
	tail := pos{0, 0}

	for _, move := range rows {
		dirAndCount := strings.Split(move, " ")
		dir := dirAndCount[0]
		count := utils.StringToInt(dirAndCount[1])
		for i := 0; i < count; i++ {
			prevHeadRow := head.row
			prevHeadCol := head.col
			head.row += dirMap[dir].row
			head.col += dirMap[dir].col
			if !areTouching(head, tail) {
				visitedPos[tail] += 1
				tail.row = prevHeadRow
				tail.col = prevHeadCol
			}
		}
	}
	return len(visitedPos) + 1
}

func areTouching(head, tail pos) bool {
	if head.row == tail.row && head.col == tail.col {
		return true
	}
	return abs(head.row-tail.row) <= 1 && abs(head.col-tail.col) <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
