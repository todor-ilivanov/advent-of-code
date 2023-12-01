package day12

import (
	"math"
	"strings"

	"advent/utils"
)

type pos struct {
	row int
	col int
}

type node struct {
	val    rune
	coords pos
}

func Solve() (int, int) {

	input := utils.ReadFileToString("2022/day12/input.txt")
	rows := strings.Split(input, "\n")

	grid, start, end := initGrid(rows)

	grid[start.row][start.col] = 'a'
	grid[end.row][end.col] = 'z'

	var startingPoints []pos

	for r, row := range grid {
		for c, el := range row {
			if el == 'a' {
				startingPoints = append(startingPoints, pos{r, c})
			}
		}
	}

	part1, _ := bfs(grid, start, end)
	minPath := math.MaxInt

	for _, startingPoint := range startingPoints {
		if curr, reachedEnd := bfs(grid, startingPoint, end); reachedEnd {
			minPath = utils.Min(curr, minPath)
		}
	}

	return part1, minPath
}

func bfs(grid [][]rune, start, end pos) (int, bool) {

	seen := make(map[pos]bool)
	seen[start] = true

	queue := make([]node, 0)
	startNode := node{grid[start.row][start.col], start}
	queue = append(queue, startNode)

	var steps int

	for len(queue) > 0 {

		var curr node

		for _, el := range queue {
			curr, queue = queue[0], queue[1:]
			row := curr.coords.row
			col := curr.coords.col

			if curr.coords == end {
				return steps, true
			}

			topPos := pos{row + 1, col}
			botPos := pos{row - 1, col}
			leftPos := pos{row, col + 1}
			rightPos := pos{row, col - 1}

			if row+1 < len(grid) && shouldExplore(seen, topPos, el.val, grid[row+1][col]) {
				neighbour := node{grid[row+1][col], topPos}
				seen[topPos] = true
				queue = append(queue, neighbour)
			}

			if row-1 >= 0 && shouldExplore(seen, botPos, el.val, grid[row-1][col]) {
				neighbour := node{grid[row-1][col], botPos}
				seen[botPos] = true
				queue = append(queue, neighbour)
			}

			if col+1 < len(grid[0]) && shouldExplore(seen, leftPos, el.val, grid[row][col+1]) {
				neighbour := node{grid[row][col+1], leftPos}
				seen[leftPos] = true
				queue = append(queue, neighbour)
			}

			if col-1 >= 0 && shouldExplore(seen, rightPos, el.val, grid[row][col-1]) {
				neighbour := node{grid[row][col-1], rightPos}
				seen[rightPos] = true
				queue = append(queue, neighbour)
			}
		}

		steps++
	}

	return steps, false
}

func shouldExplore(seen map[pos]bool, position pos, from rune, to rune) bool {
	return !isSeen(seen, position) && isReachable(from, to)
}

func isSeen(seen map[pos]bool, position pos) bool {
	if _, ok := seen[position]; ok {
		return true
	}
	return false
}

func isReachable(from rune, to rune) bool {
	return from-to >= -1
}

func initGrid(inputRows []string) ([][]rune, pos, pos) {
	grid := make([][]rune, len(inputRows))
	for i := range grid {
		grid[i] = make([]rune, len(inputRows[0]))
	}

	var start pos
	var end pos

	for r, row := range inputRows {
		for c, col := range row {
			grid[r][c] = col
			if string(col) == "S" {
				start = pos{r, c}
			} else if string(col) == "E" {
				end = pos{r, c}
			}
		}
	}
	return grid, start, end
}
