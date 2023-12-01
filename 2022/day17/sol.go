package day17

import (
	"advent/utils"
	"strconv"
	"strings"
)

type Pos struct {
	row int
	col int
}

type Rock struct {
	pebbles []Pos
	height  int
}

var dirIdx int

func Solve() (int, int) {

	input := utils.ReadFileToString("day17/input.txt")
	dirs := strings.Split(input, "")

	var grid [][]string
	grid = append(grid, make([]string, 0))

	for i := 0; i < 7; i++ {
		grid[0] = append(grid[0], "-")
	}

	part1 := dropRocks(grid, 2022, dirs)

	// part2 := dropRocks(grid, 20000, dirs)

	return part1, 0
}

func generateKey(grid [][]string, rockIdx int) string {
	topPebbleColsIdx := make([]int, 0)

	for i, el := range grid[0] {
		if el == "#" {
			topPebbleColsIdx = append(topPebbleColsIdx, i)
		}
	}

	key := []int{dirIdx, rockIdx}
	key = append(key, topPebbleColsIdx...)

	var indices []string
	for _, n := range key {
		indices = append(indices, strconv.Itoa(n))
	}

	return strings.Join(indices, "-")
}

func dropRocks(grid [][]string, numRocks int, dirs []string) int {

	rock1 := Rock{[]Pos{{0, 2}, {0, 3}, {0, 4}, {0, 5}}, 1}
	rock2 := Rock{[]Pos{{0, 3}, {1, 3}, {2, 3}, {1, 2}, {1, 4}}, 3}
	rock3 := Rock{[]Pos{{0, 4}, {1, 4}, {2, 2}, {2, 3}, {2, 4}}, 3}
	rock4 := Rock{[]Pos{{0, 2}, {1, 2}, {2, 2}, {3, 2}}, 4}
	rock5 := Rock{[]Pos{{0, 2}, {1, 2}, {0, 3}, {1, 3}}, 2}

	rocks := []Rock{rock1, rock2, rock3, rock4, rock5}

	droppedCount := 0

	for true {
		for _, rock := range rocks {
			if droppedCount == numRocks {
				return len(grid) - 1 // subtract the bottom
			}

			grid = dropRock(grid, rock, dirs)
			droppedCount++
		}
	}

	return len(grid) - 1
}

func dropRock(grid [][]string, rock Rock, dirs []string) [][]string {

	dir := getDir(dirs)
	grid = generateNewRock(grid, rock)
	grid, rock = shiftRockInDir(grid, rock, dir)
	grid = moveRock(grid, rock, dirs)

	return cleanUpGrid(grid)
}

func generateNewRock(grid [][]string, rock Rock) [][]string {
	for i := 0; i < rock.height+3; i++ {
		emptyRow := make([]string, 0)
		for i := 0; i < 7; i++ {
			emptyRow = append(emptyRow, ".")
		}
		newRow := [][]string{emptyRow}
		grid = append(newRow, grid...)
	}

	for _, pebble := range rock.pebbles {
		grid[pebble.row][pebble.col] = "@"
	}
	return grid
}

func moveRock(grid [][]string, rock Rock, dirs []string) [][]string {
	for true {

		newRockPebbles := make([]Pos, 0)

		for _, pebble := range rock.pebbles {

			newRockPebbles = append(newRockPebbles, Pos{pebble.row + 1, pebble.col})
			grid[pebble.row][pebble.col] = "."

			for _, newPebble := range newRockPebbles {
				grid[newPebble.row][newPebble.col] = "@"
			}
		}

		newRock := Rock{newRockPebbles, rock.height}

		grid, newRock = shiftRockInDir(grid, newRock, getDir(dirs))
		if hasHitBottomOrRock(grid, newRock) {
			return transformRockToResting(grid, newRock)
		}

		rock = newRock
	}
	return grid
}

func shiftRockInDir(grid [][]string, rock Rock, dir int) ([][]string, Rock) {

	for _, pebble := range rock.pebbles {
		if pebble.col+dir >= len(grid[0]) || pebble.col+dir < 0 {
			return grid, rock
		}
		if grid[pebble.row][pebble.col+dir] == "#" {
			return grid, rock
		}
	}

	newPebbles := make([]Pos, 0)

	for _, pebble := range rock.pebbles {
		newPebbles = append(newPebbles, Pos{pebble.row, pebble.col + dir})
		grid[pebble.row][pebble.col+dir] = "@"
		grid[pebble.row][pebble.col] = "."
	}

	return grid, Rock{newPebbles, rock.height}
}

func getDir(dirs []string) int {

	dirStr := dirs[dirIdx]
	dirIdx++

	if dirIdx >= len(dirs) {
		dirIdx = 0
	}

	switch dirStr {
	case ">":
		return 1
	case "<":
		return -1
	}

	return 0
}

func cleanUpGrid(grid [][]string) [][]string {
	for _, row := range grid {
		for _, el := range row {
			if el != "." {
				return grid
			}
		}
		grid = grid[1:]
	}
	return grid
}

func hasHitBottomOrRock(grid [][]string, rock Rock) bool {

	for _, pebble := range rock.pebbles {
		bottomRowIdx := pebble.row + 1
		if grid[bottomRowIdx][pebble.col] == "-" || grid[bottomRowIdx][pebble.col] == "#" {
			return true
		}
	}

	return false
}

func transformRockToResting(grid [][]string, rock Rock) [][]string {
	for _, pebble := range rock.pebbles {
		grid[pebble.row][pebble.col] = "#"
	}
	return grid
}
