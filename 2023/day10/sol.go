package day10

import (
	"advent/utils"
	"fmt"
	"math"
	"strings"
)

type pipe struct {
	first  [2]int
	second [2]int
}

func newPipe(ch rune) *pipe {
	switch ch {
	case '|':
		return &pipe{[2]int{1, 0}, [2]int{-1, 0}}
	case '-':
		return &pipe{[2]int{0, 1}, [2]int{0, -1}}
	case 'L':
		return &pipe{[2]int{-1, 0}, [2]int{0, 1}}
	case 'J':
		return &pipe{[2]int{-1, 0}, [2]int{0, -1}}
	case '7':
		return &pipe{[2]int{0, -1}, [2]int{1, 0}}
	case 'F':
		return &pipe{[2]int{0, 1}, [2]int{1, 0}}
	}
	return nil
}

var PIPES = map[rune]*pipe{
	'|': newPipe('|'),
	'-': newPipe('-'),
	'L': newPipe('L'),
	'J': newPipe('J'),
	'7': newPipe('7'),
	'F': newPipe('F'),
}

func Solve() (int, int) {
	input := utils.ReadFileToString("2023/day10/input.txt")
	rows := strings.Split(input, "\n")

	grid, startR, startC := parseGrid(rows)

	// start for input is |
	// start for sample is F
	startPipe := 'F'
	loopCoords := traverse(grid, startPipe, startR, startC)
	grid[startR][startC] = startPipe

	for r, row := range grid {
		for c := range row {
			coords := [2]int{r, c}
			if !loopCoords[coords] {
				grid[r][c] = '.'
			}
		}
	}

	// utils.PrintGrid[rune](grid)

	left, right, top, bottom := getLoopBounds(loopCoords)

	fmt.Println(left, right, top, bottom)

	return (len(loopCoords) + 1) / 2, 0
}

func parseGrid(rows []string) ([][]rune, int, int) {
	grid := make([][]rune, len(rows))
	startR, startC := 0, 0

	for r, row := range rows {
		grid[r] = make([]rune, len(row))
		for c, col := range row {
			grid[r][c] = col
			if col == 'S' {
				startR = r
				startC = c
			}
		}
	}

	return grid, startR, startC
}

func traverse(grid [][]rune, startPipe rune, startR, startC int) map[[2]int]bool {

	r := startR + PIPES[startPipe].first[0]
	c := startC + PIPES[startPipe].first[1]
	prevR, prevC := startR, startC

	loopCoords := make(map[[2]int]bool)
	loopCoords[[2]int{r, c}] = true

	for grid[r][c] != 'S' {

		current := grid[r][c]
		currentPipe := PIPES[current]

		newR := r + currentPipe.first[0]
		newC := c + currentPipe.first[1]

		if newR == prevR && newC == prevC {
			newR = r + currentPipe.second[0]
			newC = c + currentPipe.second[1]
		}

		loopCoords[[2]int{r, c}] = true
		prevR, prevC = r, c
		r, c = newR, newC
	}

	return loopCoords
}

func getLoopBounds(coords map[[2]int]bool) (int, int, int, int) {
	left, right, top, bottom := math.MaxInt, 0, math.MaxInt, 0

	for rowCol := range coords {
		r, c := rowCol[0], rowCol[1]
		left = utils.Min(r, left)
		right = utils.Max(r, right)
		top = utils.Min(c, top)
		bottom = utils.Max(c, bottom)
	}

	return left, right, top, bottom
}
