package day14

import (
	"advent/utils"
	"strings"
)

func Solve() (int, int) {
	input := utils.ReadFileToString("2023/day14/input.txt")
	rows := strings.Split(input, "\n")

	grid := parseGrid(rows)

	totalLoad := 0

	for r, row := range grid {
		for c, cell := range row {
			if cell == 'O' {
				totalLoad += calculateLoad(grid, r, c)
			}
		}
	}

	utils.PrintGrid[rune](grid)

	return totalLoad, 0
}

func parseGrid(rows []string) [][]rune {
	grid := make([][]rune, len(rows))

	for r, row := range rows {
		grid[r] = make([]rune, len(row))
		for c, ch := range row {
			grid[r][c] = ch
		}
	}

	return grid
}

func calculateLoad(grid [][]rune, r, c int) int {
	currentRow := r
	spaces := 0
	for currentRow >= 0 && grid[currentRow][c] != '#' {
		if grid[currentRow][c] == '.' {
			spaces++
		}
		currentRow--
	}
	return len(grid) - (r - spaces)
}
