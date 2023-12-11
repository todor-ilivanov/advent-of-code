package main

import (
	"fmt"
	"os"

	y22d01 "advent/2022/day01"
	y22d02 "advent/2022/day02"
	y22d03 "advent/2022/day03"
	y22d04 "advent/2022/day04"
	y22d05 "advent/2022/day05"
	y22d06 "advent/2022/day06"
	y22d07 "advent/2022/day07"
	y22d08 "advent/2022/day08"
	y22d09 "advent/2022/day09"
	y22d10 "advent/2022/day10"
	y22d11 "advent/2022/day11"
	y22d12 "advent/2022/day12"
	y22d13 "advent/2022/day13"
	y22d14 "advent/2022/day14"
	y22d15 "advent/2022/day15"
	y22d17 "advent/2022/day17"
	y22d18 "advent/2022/day18"
	y22d21 "advent/2022/day21"
	y22d23 "advent/2022/day23"

	y23d01 "advent/2023/day01"
	y23d02 "advent/2023/day02"
	y23d03 "advent/2023/day03"
	y23d04 "advent/2023/day04"
	y23d05 "advent/2023/day05"
	y23d06 "advent/2023/day06"
	y23d07 "advent/2023/day07"
	y23d08 "advent/2023/day08"
	y23d09 "advent/2023/day09"
	y23d10 "advent/2023/day10"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please specify the year and  to run (e.g., 2022 01)")
		return
	}

	year, day := os.Args[1], os.Args[2]

	switch year {
	case "2022":
		switch day {
		case "01":
			y22d01.Run()
		case "02":
			y22d02.Run()
		case "03":
			y22d03.Run()
		case "04":
			y22d04.Run()
		case "05":
			y22d05.Run()
		case "06":
			y22d06.Run()
		case "07":
			y22d07.Run()
		case "08":
			y22d08.Run()
		case "09":
			y22d09.Run()
		case "10":
			y22d10.Run()
		case "11":
			y22d11.Run()
		case "12":
			y22d12.Run()
		case "13":
			y22d13.Run()
		case "14":
			y22d14.Run()
		case "15":
			y22d15.Run()
		case "17":
			y22d17.Run()
		case "18":
			y22d18.Run()
		case "21":
			y22d21.Run()
		case "23":
			y22d23.Run()
		}
	case "2023":
		switch day {
		case "01":
			y23d01.Run()
		case "02":
			y23d02.Run()
		case "03":
			y23d03.Run()
		case "04":
			y23d04.Run()
		case "05":
			y23d05.Run()
		case "06":
			y23d06.Run()
		case "07":
			y23d07.Run()
		case "08":
			y23d08.Run()
		case "09":
			y23d09.Run()
		case "10":
			y23d10.Run()
		}
	default:
		fmt.Println("Year not recognized")
	}
}
