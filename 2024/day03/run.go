package day03

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println("----Day 03----")

	start := time.Now()
	part1, part2 := Solve()
	fmt.Println(part1)
	fmt.Println(part2)

	fmt.Println("Time elapsed:", time.Since(start))
}
