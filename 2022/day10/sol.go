package day10

import (
	"fmt"
	"strings"

	"advent/utils"
)

var crt [40 * 6]string

func Solve() int {

	input := utils.ReadFileToString("2022/day10/input.txt")
	rows := strings.Split(input, "\n")

	cycle := 0
	x := 1
	var result int

	q := make(map[int]int)

	for _, command := range rows {
		cycle++
		if !(command == "noop") {
			val := strings.Split(command, " ")[1]
			q[cycle+1] = utils.StringToInt(val)
			result += getCycleResult(cycle, x)
			renderOnCrt(cycle, x, cycle/40)
			cycle++
		}
		result += getCycleResult(cycle, x)
		renderOnCrt(cycle, x, cycle/40)
		x += q[cycle]
	}

	for i := 0; i < len(crt); i += 40 {
		fmt.Println(crt[i : i+40])
	}

	return result
}

func renderOnCrt(cycle, x, row int) {
	pixel := cycle - 1
	pos := row*40 + x
	if pixel == pos || pixel == pos+1 || pixel == pos-1 {
		crt[pixel] = "#"
	} else {
		crt[pixel] = "."
	}
}

func getCycleResult(cycle, x int) int {
	if cycle == 20 || ((cycle-20)%40 == 0) {
		return cycle * x
	}
	return 0
}
