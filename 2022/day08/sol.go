package day08

import (
	"strings"

	"advent/utils"
)

type node struct {
	top   int
	bot   int
	left  int
	right int
}

func SolvePart1() int {
	input := utils.ReadFileToString("2022/day08/input.txt")
	rows := strings.Split(input, "\n")

	grid, dp := initGrids(rows)

	result := len(grid)*2 + len(grid[0])*2 - 4

	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			curr := grid[r][c]
			dp[r][c].left = utils.Max(dp[r][c-1].left, curr)
			dp[r][c].top = utils.Max(dp[r-1][c].top, curr)
			dp[r][c].right = utils.Max(calcMaxRight(r, c+1, dp, grid), curr)
			dp[r][c].bot = utils.Max(calcMaxBot(r+1, c, dp, grid), curr)

			if isVisible(curr, r, c, dp) {
				result += 1
			}
		}
	}

	return result
}

func calcMaxRight(r, c int, dp [][]node, grid [][]int) int {
	if dp[r][c].right != -1 {
		return dp[r][c].right
	}
	dp[r][c].right = utils.Max(grid[r][c], calcMaxRight(r, c+1, dp, grid))
	return dp[r][c].right
}

func calcMaxBot(r, c int, dp [][]node, grid [][]int) int {
	if dp[r][c].bot != -1 {
		return dp[r][c].bot
	}
	dp[r][c].bot = utils.Max(grid[r][c], calcMaxBot(r+1, c, dp, grid))
	return dp[r][c].bot
}

func isVisible(tree, r, c int, dp [][]node) bool {
	return tree > dp[r][c-1].left || tree > dp[r-1][c].top || tree > dp[r+1][c].bot || tree > dp[r][c+1].right
}

func initGrids(rows []string) ([][]int, [][]node) {

	grid := make([][]int, len(rows))
	dp := make([][]node, len(rows))

	for i := range grid {
		grid[i] = make([]int, len(rows[0]))
		dp[i] = make([]node, len(rows[0]))
	}

	for i := 0; i < len(rows); i++ {
		nums := strings.Split(rows[i], "")
		for j := 0; j < len(nums); j++ {
			grid[i][j] = utils.StringToInt(nums[j])
			dp[i][j] = node{-1, -1, -1, -1}
		}
	}

	for i := 0; i < len(rows[0]); i++ {
		dp[0][i].top = grid[0][i]
	}

	for i := 0; i < len(rows); i++ {
		dp[i][0].left = grid[i][0]
	}

	for i := len(rows) - 1; i >= 0; i-- {
		dp[i][len(rows)-1].right = grid[i][len(rows)-1]
	}

	for i := len(rows[0]) - 1; i >= 0; i-- {
		dp[len(rows[0])-1][i].bot = grid[len(rows[0])-1][i]
	}

	return grid, dp
}
