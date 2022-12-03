package day08

import (
	"sort"
	"strings"

	"advent/utils"
)

// solution is embarrasing... uplaoding it anyway
func SolvePart2() int {
	input := utils.ReadFileToString("day08/input.txt")
	rows := strings.Split(input, "\n")

	grid, visMap := initGridsPart2(rows)

	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			visMap[r][c].left = calcVisLeft(r, c, 0, grid)
			visMap[r][c].top = calcVisTop(r, c, 0, grid)
			visMap[r][c].right = calcVisRight(r, c, 0, grid)
			visMap[r][c].bot = calcVisBot(r, c, 0, grid)
		}
	}

	var scores []int

	for _, r := range visMap {
		for _, c := range r {
			result := c.bot * c.top * c.left * c.right
			scores = append(scores, result)
		}
	}

	sort.Ints(scores)

	return scores[len(scores)-1]
}

func calcVisRight(r, c, sum int, grid [][]int) int {

	colToRight := c + 1
	for colToRight <= len(grid[0])-1 {
		if grid[r][c] <= grid[r][colToRight] {
			return sum + 1
		}
		sum++
		colToRight++
	}
	return sum
}

func calcVisLeft(r, c, sum int, grid [][]int) int {

	colToLeft := c - 1
	for colToLeft >= 0 {
		if grid[r][c] <= grid[r][colToLeft] {
			return sum + 1
		}
		sum++
		colToLeft--
	}
	return sum
}

func calcVisTop(r, c, sum int, grid [][]int) int {

	rowAbove := r - 1
	for rowAbove >= 0 {
		if grid[r][c] <= grid[rowAbove][c] {
			return sum + 1
		}
		sum++
		rowAbove--
	}
	return sum
}

func calcVisBot(r, c, sum int, grid [][]int) int {

	rowBelow := r + 1
	for rowBelow <= len(grid)-1 {
		if grid[r][c] <= grid[rowBelow][c] {
			return sum + 1
		}
		sum++
		rowBelow++
	}
	return sum
}

func initGridsPart2(rows []string) ([][]int, [][]node) {

	grid := make([][]int, len(rows))
	visMap := make([][]node, len(rows))

	for i := range grid {
		grid[i] = make([]int, len(rows[0]))
		visMap[i] = make([]node, len(rows[0]))
	}

	for i := 0; i < len(rows); i++ {
		nums := strings.Split(rows[i], "")
		for j := 0; j < len(nums); j++ {
			grid[i][j] = utils.StringToInt(nums[j])
			visMap[i][j] = node{-1, -1, -1, -1}
		}
	}

	return grid, visMap
}
