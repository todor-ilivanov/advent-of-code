package day01

import (
	"advent/utils"
	"sort"
	"strconv"
	"strings"
)

func Solve() (int, int) {
	input := utils.ReadFileToString("2024/day01/input.txt")
	rows := strings.Split(input, "\n")

	var col1 []int
	var col2 []int
	freqs := make(map[int]int)

	for _, row := range rows {
		r := strings.Fields(row)
		firstNum, _ := strconv.Atoi(r[0])
		secondNum, _ := strconv.Atoi(r[1])

		freqs[secondNum]++

		col1 = append(col1, firstNum)
		col2 = append(col2, secondNum)
	}

	sort.Ints(col1)
	sort.Ints(col2)

	part1 := 0
	for i := 0; i < len(col1); i++ {
		part1 += utils.Abs(col1[i] - col2[i])
	}

	part2 := 0
	for _, num := range col1 {
		part2 += num * freqs[num]
	}

	return part1, part2
}
