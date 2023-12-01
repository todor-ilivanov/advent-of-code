package day01

import (
	"fmt"
	"strconv"
	"strings"

	"advent/utils"
)

func SolvePart1() int {

	str := utils.ReadFileToString("2022/day01/input.txt")

	allReadings := strings.Split(str, "\n\n")

	var res int

	for _, el := range allReadings {
		var readings = strings.Split(el, "\n")
		var sum int
		for _, el := range readings {
			num, err := strconv.Atoi(el)
			if err != nil {
				fmt.Print(err)
			}
			sum += num
		}
		res = utils.Max(res, sum)
	}

	return res
}
