package day15

import (
	"regexp"
	"strings"

	"advent/utils"
)

type pos struct {
	row int
	col int
}

type interval struct {
	start int
	end   int
}

var Y int = 2000000

func Solve() (int, int) {

	input := utils.ReadFileToString("day15/input.txt")
	rows := strings.Split(input, "\n")

	readingMap := make(map[pos]string)
	for _, reading := range rows {
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(reading, -1)
		sensorX := utils.StringToInt(nums[0])
		sensorY := utils.StringToInt(nums[1])
		beaconX := utils.StringToInt(nums[2])
		beaconY := utils.StringToInt(nums[3])

		sensor := pos{sensorY, sensorX}
		beacon := pos{beaconY, beaconX}

		readingMap[sensor] = "S"
		readingMap[beacon] = "B"

		dist := calcDistance(sensor, beacon)

	}

	// fmt.Println(readingMap)

	var res int
	for coords, el := range readingMap {
		if coords.row == 2000000 && el == "#" {
			res++
		}
	}

	return res, 0
}

func calcDistance(sensor, beacon pos) int {
	return utils.Abs(sensor.row-beacon.row) + utils.Abs(sensor.col-beacon.col)
}

func drawRange(readingMap map[pos]string, sensor pos, dist int) map[pos]string {

	startRow := sensor.row - dist

	toCoverEitherSide := 0

	for row := startRow; row < sensor.row; row++ {

		// if row == 2000000 {
		drawOnRow(readingMap, pos{row, sensor.col}, toCoverEitherSide)
		// }

		toCoverEitherSide++
	}

	for row := sensor.row; row <= sensor.row+dist; row++ {

		// if row == 2000000 {
		drawOnRow(readingMap, pos{row, sensor.col}, toCoverEitherSide)
		// }

		toCoverEitherSide--
	}

	return readingMap
}

func drawOnRow(readingMap map[pos]string, position pos, colOffset int) map[pos]string {

	for i := 0; i <= colOffset; i++ {
		currLeft := pos{position.row, position.col + i}
		currRight := pos{position.row, position.col - i}

		readingMap = drawOnPos(readingMap, currLeft)
		readingMap = drawOnPos(readingMap, currRight)
	}

	return readingMap
}

func drawOnPos(readingMap map[pos]string, position pos) map[pos]string {

	if !isSensorOrBeacon(readingMap, position) {
		readingMap[position] = "#"
	}
	return readingMap
}

func isSensorOrBeacon(readingMap map[pos]string, el pos) bool {
	return readingMap[el] == "S" || readingMap[el] == "B"
}
