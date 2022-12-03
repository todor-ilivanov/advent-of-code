package main

import (
	"fmt"

	"advent/day01"
	"advent/day02"
	"advent/day03"
	"advent/day04"
	"advent/day05"
	"advent/day06"
	"advent/day07"
	"advent/day08"
	"advent/day09"
)

func main() {
	// solveDay01()
	// solveDay02()
	// solveDay03()
	// solveDay04()
	// solveDay05()
	// solveDay06()
	// solveDay07()
	// solveDay08()
	solveDay09()
}

func solveDay01() {
	fmt.Println("----Day 01----")
	part1 := day01.SolvePart1()
	part2 := day01.SolvePart2()
	fmt.Println(part1)
	fmt.Println(part2)
}

func solveDay02() {
	fmt.Println("----Day 02----")
	part1 := day02.SolvePart1()
	part2 := day02.SolvePart2()
	fmt.Println(part1)
	fmt.Println(part2)
}

func solveDay03() {
	fmt.Println("----Day 03----")
	part1 := day03.SolvePart1()
	part2 := day03.SolvePart2()
	fmt.Println(part1)
	fmt.Println(part2)
}

func solveDay04() {
	fmt.Println("----Day 04----")
	part1, part2 := day04.Solve()
	fmt.Println(part1)
	fmt.Println(part2)
}

func solveDay05() {
	fmt.Println("----Day 05----")
	part1 := day05.SolvePart1()
	part2 := day05.SolvePart2()
	fmt.Println(part1)
	fmt.Println(part2)
}

func solveDay06() {
	fmt.Println("----Day 06----")
	part1 := day06.SolveForWindowSize(4)
	part2 := day06.SolveForWindowSize(14)
	fmt.Println(part1)
	fmt.Println(part2)
}

func solveDay07() {
	fmt.Println("----Day 07----")
	part1, part2 := day07.Solve()
	fmt.Println(part1)
	fmt.Println(part2)
}

func solveDay08() {
	fmt.Println("----Day 08----")
	part1 := day08.SolvePart1()
	fmt.Println(part1)
	part2 := day08.SolvePart2()
	fmt.Println(part2)
}

func solveDay09() {
	fmt.Println("----Day 09----")
	part1 := day09.SolvePart1()
	fmt.Println(part1)
	part2 := day09.SolvePart2(10)
	fmt.Println(part2)
}
