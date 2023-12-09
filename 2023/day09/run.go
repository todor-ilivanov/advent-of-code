package day09

import (
	"advent/utils"
	"fmt"
	"strings"
	"time"
)

func Run() {
	fmt.Println("----Day 09----")

	input := utils.ReadFileToString("2023/day09/input.txt")
	rows := strings.Split(input, "\n")

	start := time.Now()
	part1 := SolvePart1(rows)
	fmt.Printf("%-15d Time: %s\n", part1, time.Since(start))

	start = time.Now()
	part2 := SolvePart2(rows)
	fmt.Printf("%-15d Time: %s\n", part2, time.Since(start))

}
