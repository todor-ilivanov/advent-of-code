package day04

import (
	"advent/utils"
	"fmt"
)

func Solve() (int, int) {
	input := utils.ReadFileToString("2024/day04/input.txt")

	grid := utils.ParseInputToGrid(input)

	return solvePart1(grid), solvePart2(grid)
}

func solvePart1(grid [][]rune) int {

	rows := len(grid)
	cols := len(grid[0])

	XMAS := "XMAS"
	SAMX := "SAMX"

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

	return count
}

func solvePart2(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])

	windowSize := 3

	count := 0

	for r := 0; r <= rows-windowSize; r++ {
		for c := 0; c <= cols-windowSize; c++ {
			subgrid := getSubgrid(grid, windowSize, r, c)
			if checkSubgrid(subgrid, windowSize) {
				count++
			}
			fmt.Println("------")
		}
	}

	return count
}

func getSubgrid(grid [][]rune, size int, r int, c int) [][]rune {
	subgrid := make([][]rune, size)
	for i := range subgrid {
		subgrid[i] = make([]rune, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			subgrid[i][j] = grid[r+i][c+j]
			fmt.Printf("%c ", subgrid[i][j])
		}
		fmt.Println()
	}

	return subgrid
}

func checkSubgrid(grid [][]rune, size int) bool {

	MAS := "MAS"
	SAM := "SAM"

	var diag1 string
	for i := 0; i < size; i++ {
		diag1 += string(grid[i][i])
	}

	var diag2 string
	for i := 0; i < size; i++ {
		row := size - 1 - i
		diag2 += string(grid[row][i])
	}

	fmt.Println(diag1)
	fmt.Println(diag2)
	return (diag1 == MAS || diag1 == SAM) && (diag2 == MAS || diag2 == SAM)
}
