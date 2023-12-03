package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func ReadFileToString(name string) string {
	b, err := os.ReadFile(name)

	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Print(err)
	}
	return num
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(num int) int {
	numFloat := float64(num)
	return int(math.Abs(numFloat))
}

func Trim(str string) string {
	return str[:len(str)-1]
}

func PrintGrid[T any](grid [][]T) {
	for _, row := range grid {
		for _, col := range row {
			switch v := any(col).(type) {
			case rune:
				fmt.Print(fmt.Sprintf("%c ", v))
			default:
				fmt.Print(fmt.Sprintf("%v ", v))
			}
		}
		fmt.Println()
	}
}

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Multiply(nums []int) int {
	sum := 1
	for _, num := range nums {
		sum *= num
	}
	return sum
}
