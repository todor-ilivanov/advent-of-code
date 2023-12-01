package day09

import (
	"strings"

	"advent/utils"
)

type knot struct {
	tail   *knot
	coords pos
}

func SolvePart2(numKnots int) int {
	input := utils.ReadFileToString("2022/day09/input.txt")
	rows := strings.Split(input, "\n")

	dirMap := make(map[string]pos)
	dirMap["R"] = pos{0, 1}
	dirMap["L"] = pos{0, -1}
	dirMap["U"] = pos{-1, 0}
	dirMap["D"] = pos{1, 0}
	dirMap["UR"] = pos{-1, 1}
	dirMap["UL"] = pos{-1, -1}
	dirMap["DR"] = pos{1, 1}
	dirMap["DL"] = pos{1, -1}

	visitedPos := make(map[pos]int)
	knots := make([]knot, numKnots)

	for i := 0; i < numKnots-1; i++ {
		knots[i].tail = &knots[i+1]
	}

	for _, move := range rows {
		dirAndCount := strings.Split(move, " ")
		dir := dirAndCount[0]
		count := utils.StringToInt(dirAndCount[1])

		for i := 0; i < count; i++ {
			knots[0].coords.row += dirMap[dir].row
			knots[0].coords.col += dirMap[dir].col

			// fmt.Println("Moved head in", dir, "from", headOld, "to", knots[0].coords)

			moveTail(&knots[0], dirMap)
			lastNodeCoords := knots[len(knots)-1].coords
			visitedPos[lastNodeCoords] += 1
		}
	}

	return len(visitedPos)
}

func moveTail(knot *knot, dirMap map[string]pos) {
	if knot.tail == nil || areTouching(knot.coords, knot.tail.coords) {
		return
	}

	dir := getTailDir(knot)

	knot.tail.coords.row += dirMap[dir].row
	knot.tail.coords.col += dirMap[dir].col
	// fmt.Println("Moved", knot.tail, "in", dir, "from", coordsOld, "to", knot.tail.coords)
	moveTail(knot.tail, dirMap)
}

func getTailDir(knot *knot) string {

	if knot == nil || knot.tail == nil {
		return ""
	}

	if knot.coords.col > knot.tail.coords.col && knot.coords.row == knot.tail.coords.row {
		return "R"
	} else if knot.coords.col < knot.tail.coords.col && knot.coords.row == knot.tail.coords.row {
		return "L"
	} else if knot.coords.col == knot.tail.coords.col && knot.coords.row < knot.tail.coords.row {
		return "U"
	} else if knot.coords.col == knot.tail.coords.col && knot.coords.row > knot.tail.coords.row {
		return "D"
	} else if knot.coords.col > knot.tail.coords.col && knot.coords.row < knot.tail.coords.row {
		return "UR"
	} else if knot.coords.col < knot.tail.coords.col && knot.coords.row < knot.tail.coords.row {
		return "UL"
	} else if knot.coords.col > knot.tail.coords.col && knot.coords.row > knot.tail.coords.row {
		return "DR"
	} else if knot.coords.col < knot.tail.coords.col && knot.coords.row > knot.tail.coords.row {
		return "DL"
	}

	return ""
}
