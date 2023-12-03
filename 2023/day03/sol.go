package day03

import (
	"advent/utils"
	"math"
	"strings"
	"unicode"
)

type dir struct {
	row int
	col int
}

var SYMBOLS = map[rune]bool{
	'#': true,
	'$': true,
	'%': true,
	'&': true,
	'*': true,
	'+': true,
	'-': true,
	'/': true,
	'=': true,
	'@': true,
}

func Solve() (int, int) {
	input := utils.ReadFileToString("2023/day03/input.txt")

	grid := parseInputToGrid(input)
	part1Sum := 0
	part2Sum := 0

	for r := range grid {
		for c := range grid[r] {

			cell := grid[r][c]

			if isSymbol(cell) {
				updatedGrid, adjacentNumbers := sweepAdjacentNumbers(grid, r, c)

				part1Sum += utils.Sum(adjacentNumbers)

				if cell == '*' && len(adjacentNumbers) == 2 {
					part2Sum += utils.Multiply(adjacentNumbers)
				}

				grid = updatedGrid
			}
		}
	}

	return part1Sum, part2Sum
}

func parseInputToGrid(input string) [][]rune {
	rows := strings.Split(input, "\n")

	ROWS := len(rows)
	COLS := len(rows[0])

	grid := make([][]rune, ROWS)

	for i := 0; i < ROWS; i++ {
		grid[i] = make([]rune, COLS)
	}

	for r, row := range rows {
		for c, col := range row {
			grid[r][c] = col
		}
	}

	return grid
}

func isSymbol(ch rune) bool {
	_, ok := SYMBOLS[ch]
	return ok
}

func sweepAdjacentNumbers(grid [][]rune, r int, c int) ([][]rune, []int) {
	dirs := []dir{{0, 1}, {0, -1}, {1, 0}, {1, -1}, {-1, 1}, {-1, 0}, {1, 1}, {-1, -1}}

	adjacentNumbers := make([]int, 0)

	for _, dir := range dirs {
		newR, newC := r+dir.row, c+dir.col

		if newR < 0 || newR > len(grid) || newC < 0 || newC > len(grid[newR]) {
			continue
		}

		if unicode.IsDigit(grid[newR][newC]) {
			updatedRow, number := updateRowAndConstructNumber(grid[newR], newC)
			adjacentNumbers = append(adjacentNumbers, number)
			grid[newR] = updatedRow
		}
	}
	return grid, adjacentNumbers
}

func updateRowAndConstructNumber(row []rune, idx int) ([]rune, int) {
	l, r := idx-1, idx+1
	digits := []int{int(row[idx] - '0')}
	row[idx] = '.'

	for l >= 0 && row[l] != '.' && !isSymbol(row[l]) {
		digit := int(row[l] - '0')
		digits = append([]int{digit}, digits...)
		row[l] = '.'
		l -= 1
	}

	for r < len(row) && row[r] != '.' && !isSymbol(row[r]) {
		digit := int(row[r] - '0')
		digits = append(digits, digit)
		row[r] = '.'
		r += 1
	}

	return row, constructNumber(digits)
}

func constructNumber(digits []int) int {
	constructedNum := 0
	for i := 0; i < len(digits); i++ {
		pow := len(digits) - i - 1
		constructedNum += digits[i] * int(math.Pow10(pow))
	}
	return constructedNum
}
