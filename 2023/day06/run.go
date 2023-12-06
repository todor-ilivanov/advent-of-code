package day06

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println("----Day 06----")

	start := time.Now()
	part1 := SolvePart1()
	fmt.Println(part1, "Time elapsed:", time.Since(start))

	start = time.Now()
	part2 := SolvePart2()
	fmt.Println(part2, "Time elapsed:", time.Since(start))
}
