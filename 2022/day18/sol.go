package day18

import (
	"strings"

	"advent/utils"
)

type Pos struct {
	x int
	y int
	z int
}

type Cube struct {
	pos             Pos
	maybeNeighbours []Pos
}

const CUBE_SIDES int = 6

func Solve() (int, int) {

	input := utils.ReadFileToString("2022/day18/input.txt")
	rows := strings.Split(input, "\n")

	cubes := make([]Cube, 0)
	lava := make(map[Pos]bool, 0)

	for _, row := range rows {
		dims := strings.Split(row, ",")

		x := utils.StringToInt(dims[0])
		y := utils.StringToInt(dims[1])
		z := utils.StringToInt(dims[2])

		cubePos := Pos{x, y, z}
		lava[cubePos] = true
		cubes = append(cubes, Cube{cubePos, calcPotentialNeighbours(cubePos)})
	}

	part1 := 0

	for _, cube := range cubes {
		neighbours := getNeighbours(lava, cube)
		part1 += CUBE_SIDES - len(neighbours)
	}

	return part1, 0
}

func getNeighbours(lava map[Pos]bool, cube Cube) []Pos {
	neighbours := make([]Pos, 0)
	for _, maybeNeighbour := range cube.maybeNeighbours {
		if contains(lava, maybeNeighbour) {
			neighbours = append(neighbours, maybeNeighbour)
		}
	}
	return neighbours
}

func calcPotentialNeighbours(pos Pos) []Pos {

	maybeNeighbours := make([]Pos, 0)
	x, y, z := pos.x, pos.y, pos.z

	maybeNeighbours = append(maybeNeighbours, Pos{x + 1, y, z})
	maybeNeighbours = append(maybeNeighbours, Pos{x - 1, y, z})
	maybeNeighbours = append(maybeNeighbours, Pos{x, y + 1, z})
	maybeNeighbours = append(maybeNeighbours, Pos{x, y - 1, z})
	maybeNeighbours = append(maybeNeighbours, Pos{x, y, z + 1})
	maybeNeighbours = append(maybeNeighbours, Pos{x, y, z - 1})

	return maybeNeighbours
}

func contains(posMap map[Pos]bool, pos Pos) bool {
	if _, ok := posMap[pos]; ok {
		return ok
	}
	return false
}
