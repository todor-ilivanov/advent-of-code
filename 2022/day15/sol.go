package day15

import (
	"regexp"
	"strings"

	"advent/utils"
)

type pos struct {
	x int
	y int
}

type interval struct {
	start int
	end   int
}

var Y int = 2000000

func Solve() (int, int) {

	input := utils.ReadFileToString("day15/input.txt")
	rows := strings.Split(input, "\n")

	var beaconsAtY []int
	var intervals []interval

	sensorDistMap := make(map[pos]int)

	for _, reading := range rows {
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(reading, -1)
		sensorX := utils.StringToInt(nums[0])
		sensorY := utils.StringToInt(nums[1])
		beaconX := utils.StringToInt(nums[2])
		beaconY := utils.StringToInt(nums[3])

		sensor := pos{sensorX, sensorY}
		beacon := pos{beaconX, beaconY}

		dist := calcDistance(sensor, beacon)
		sensorDistMap[sensor] = dist
		if beaconY == Y && !contains(beaconsAtY, beaconX) {
			beaconsAtY = append(beaconsAtY, beaconX)
		}
	}

	var res, minX, maxX int

	for sensor, dist := range sensorDistMap {

		dx := dist - utils.Abs(sensor.y-Y)

		if dx > 0 {
			minX = utils.Min(minX, sensor.x-dx)
			maxX = utils.Max(maxX, sensor.x+dx)
			intervals = append(intervals, interval{sensor.x - dx, sensor.x + dx})
		}
	}

	for x := minX; x <= maxX; x++ {

		if contains(beaconsAtY, x) {
			continue
		}

		for _, interval := range intervals {
			if interval.start <= x && x <= interval.end {
				res++
				break
			}
		}
	}

	return res, 0
}

func contains(slice []int, el int) bool {
	for _, val := range slice {
		if val == el {
			return true
		}
	}
	return false
}

func calcDistance(sensor, beacon pos) int {
	return utils.Abs(sensor.x-beacon.x) + utils.Abs(sensor.y-beacon.y)
}
