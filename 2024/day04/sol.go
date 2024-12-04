package day04

import (
	"advent/utils"
)

func Solve() (int, int) {
	input := utils.ReadFileToString("2024/day04/input.txt")

	XMAS := "XMAS"
	SAMX := "SAMX"

	grid := utils.ParseInputToGrid(input)
	rows := len(grid)
	cols := len(grid[0])

	utils.PrintGrid(grid)

	windowSize := 4
	count := 0

	// horizontal
	for r := 0; r < rows; r++ {
		for c := 0; c <= cols-windowSize; c++ {
			word := string(grid[r][c : windowSize+c])

			if word == XMAS || word == SAMX {
				count++
			}
		}
	}

	// vertical
	for c := 0; c < cols; c++ {
		for r := 0; r <= rows-windowSize; r++ {
			word := ""
			for i := 0; i < windowSize; i++ {
				word += string(grid[r+i][c])
			}

			if word == XMAS || word == SAMX {
				count++
			}
		}
	}

	// diagonal
	for r := 0; r <= rows-windowSize; r++ {
		for c := 0; c <= cols-windowSize; c++ {
			word := ""
			for i := 0; i < windowSize; i++ {
				word += string(grid[r+i][c+i])
			}

			if word == XMAS || word == SAMX {
				count++
			}
		}
	}

	for r := 0; r <= rows-windowSize; r++ {
		for c := windowSize - 1; c < cols; c++ {
			word := ""
			for i := 0; i < windowSize; i++ {
				word += string(grid[r+i][c-i])
			}

			if word == XMAS || word == SAMX {
				count++
			}
		}
	}

	return count, 0
}
