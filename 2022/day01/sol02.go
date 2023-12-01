package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"advent/utils"
)

func SolvePart2() int {

	str := utils.ReadFileToString("day01/input.txt")

	allReadings := strings.Split(str, "\n\n")
	readingsArr := []int{}

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
		readingsArr = append(readingsArr, sum)
	}

	sort.Ints(readingsArr)
	var sumTop3 int
	for i := len(readingsArr) - 3; i < len(readingsArr); i++ {
		sumTop3 += readingsArr[i]
	}

	return sumTop3
}
