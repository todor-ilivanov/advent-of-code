package utils

import (
	"fmt"
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
