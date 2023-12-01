package day06

import (
	"fmt"
)

func Run() {
	fmt.Println("----Day 06----")
	part1, part2 := SolveForWindowSize(4), SolveForWindowSize(14)
	fmt.Println(part1)
	fmt.Println(part2)
}
