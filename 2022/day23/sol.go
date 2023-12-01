package day23

import (
	"math"
	"strings"

	"advent/utils"
)

type Pos struct {
	row int
	col int
}

func Solve(rounds int) (int, int) {

	input := utils.ReadFileToString("2022/day23/input.txt")
	rows := strings.Split(input, "\n")

	elves := initElves(rows)

	dirs := []rune{'N', 'S', 'W', 'E'}

	for i := 0; i < rounds; i++ {

		elves = playRound(elves, dirs)
		frontDir := dirs[0]
		dirs = dirs[1:]
		dirs = append(dirs, frontDir)
	}

	var minRow, minCol int = math.MaxInt, math.MaxInt
	var maxRow, maxCol int = math.MinInt, math.MinInt
	for pos := range elves {
		minRow = utils.Min(minRow, pos.row)
		maxRow = utils.Max(maxRow, pos.row)
		minCol = utils.Min(minCol, pos.col)
		maxCol = utils.Max(maxCol, pos.col)
	}

	area := (maxRow - minRow) * (maxCol - minCol)
	ans := area - len(elves)

	return ans, 0
}

func initElves(rows []string) map[Pos]bool {
	elves := make(map[Pos]bool)

	for rowIdx, row := range rows {
		for colIdx, ch := range row {
			if ch != '#' {
				continue
			}
			pos := Pos{rowIdx, colIdx}
			elves[pos] = true
		}
	}
	return elves
}

func playRound(elves map[Pos]bool, dirs []rune) map[Pos]bool {

	proposedPos := planMoves(elves, dirs)

	proposedPos = deduplicateProposedPos(proposedPos)

	movedElves := moveElves(elves, proposedPos)

	return movedElves
}

func planMoves(elves map[Pos]bool, dirs []rune) map[Pos]Pos {

	proposedPos := make(map[Pos]Pos)

	for pos := range elves {
		for _, dir := range dirs {
			dirsToCheck := dirsToCheck(dir)
			if canMoveToDir(elves, pos, dirsToCheck) {
				newPos := Pos{pos.row + dirsToCheck[1].row, pos.col + dirsToCheck[1].col}
				proposedPos[pos] = newPos
				break
			}
		}
	}

	return proposedPos
}

func moveElves(elves map[Pos]bool, proposedPos map[Pos]Pos) map[Pos]bool {
	var posToDelete []Pos

	for prevPos, newPos := range proposedPos {
		if shouldMove(elves, prevPos) {
			posToDelete = append(posToDelete, prevPos)
			elves[newPos] = true
		}
	}

	for _, pos := range posToDelete {
		delete(elves, pos)
	}
	return elves
}

func shouldMove(elves map[Pos]bool, pos Pos) bool {
	allDirs := []Pos{{-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}}
	for _, dir := range allDirs {
		neighbourPos := Pos{pos.row + dir.row, pos.col + dir.col}
		if contains(elves, neighbourPos) {
			return true
		}
	}
	return false
}

func dirsToCheck(dir rune) []Pos {
	switch dir {
	case 'N':
		return []Pos{{-1, -1}, {-1, 0}, {-1, 1}}
	case 'S':
		return []Pos{{1, -1}, {1, 0}, {1, 1}}
	case 'W':
		return []Pos{{-1, -1}, {0, -1}, {1, -1}}
	case 'E':
		return []Pos{{-1, 1}, {0, 1}, {1, 1}}
	default:
		panic("No such dir.")
	}
}

func canMoveToDir(elves map[Pos]bool, pos Pos, dirsToCheck []Pos) bool {
	for _, dir := range dirsToCheck {
		newPos := Pos{pos.row + dir.row, pos.col + dir.col}
		if contains(elves, newPos) {
			return false
		}
	}
	return true
}

func deduplicateProposedPos(proposedPos map[Pos]Pos) map[Pos]Pos {
	seen := make(map[Pos]int)

	for _, pos := range proposedPos {
		seen[pos] += 1
	}

	for prevPos, newPos := range proposedPos {
		if seen[newPos] > 1 {
			delete(proposedPos, prevPos)
		}
	}

	return proposedPos
}

func contains(posMap map[Pos]bool, pos Pos) bool {
	if _, ok := posMap[pos]; ok {
		return ok
	}
	return false
}
