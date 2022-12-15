package day14

import (
	"strings"

	"advent/utils"
)

type pos struct {
	row int
	col int
}

func Solve() (int, int) {

	input := utils.ReadFileToString("day14/input.txt")
	rows := strings.Split(input, "\n")

	caveMap := make(map[pos]string)
	var maxRockRow int

	for _, rock := range rows {

		points := strings.Split(rock, " -> ")
		prev := points[0]

		for _, curr := range points[1:] {
			prevPos := stringToPos(prev)
			currPos := stringToPos(curr)
			caveMap = drawLine(caveMap, prevPos, currPos)

			maxRockRow = utils.Max(maxRockRow, prevPos.row)
			prev = curr
		}
	}

	floorRow := maxRockRow + 2

	part1 := simulateSand(caveMap, floorRow, func(a pos) bool { return a.row >= maxRockRow })
	part2 := simulateSand(caveMap, floorRow, func(a pos) bool { return a == pos{0, 500} })

	return part1, part2
}

func simulateSand(caveMap map[pos]string, floorRow int, check func(pos) bool) int {
	var result int

	for true {
		sandPos := pos{0, 500}
		for true {
			newPos := moveSand(caveMap, sandPos)
			if check(newPos) {
				return result
			}
			if newPos == sandPos || newPos.row == floorRow {
				break
			}
			sandPos = newPos
		}
		caveMap[sandPos] = "o"
		result++
	}

	return result
}

func stringToPos(point string) pos {
	rowAndCol := strings.Split(point, ",")
	row := utils.StringToInt(rowAndCol[1])
	col := utils.StringToInt(rowAndCol[0])
	return pos{row, col}
}

func drawLine(caveMap map[pos]string, prev, curr pos) map[pos]string {
	dir := getDir(prev, curr)
	caveMap[prev] = "#"
	caveMap[curr] = "#"

	if dir.row != 0 {
		for i := 0; i < utils.Abs(prev.row-curr.row); i++ {
			newPos := pos{prev.row + (dir.row * i), prev.col}
			caveMap[newPos] = "#"
		}
	}

	if dir.col != 0 {
		for i := 0; i < utils.Abs(prev.col-curr.col); i++ {
			newPos := pos{prev.row, prev.col + (dir.col * i)}
			caveMap[newPos] = "#"
		}
	}

	return caveMap
}

func moveSand(caveMap map[pos]string, sandPos pos) pos {
	posDown := pos{sandPos.row + 1, sandPos.col}
	posDownLeft := pos{sandPos.row + 1, sandPos.col - 1}
	posDownRight := pos{sandPos.row + 1, sandPos.col + 1}

	if !isRockOrSand(caveMap[posDown]) {
		return pos{sandPos.row + 1, sandPos.col}
	}

	if !isRockOrSand(caveMap[posDownLeft]) {
		return pos{sandPos.row + 1, sandPos.col - 1}
	}

	if !isRockOrSand(caveMap[posDownRight]) {
		return pos{sandPos.row + 1, sandPos.col + 1}
	}

	return sandPos
}

func isRockOrSand(el string) bool {
	return el == "#" || el == "o"
}

func getDir(p1, p2 pos) pos {
	if p1.row == p2.row {
		if p1.col > p2.col {
			return pos{0, -1}
		} else {
			return pos{0, 1}
		}
	}

	if p1.col == p2.col {
		if p1.row > p2.row {
			return pos{-1, 0}
		} else {
			return pos{1, 0}
		}
	}

	return pos{0, 0}
}
