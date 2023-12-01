package day13

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"

	"advent/utils"
)

func Solve() (int, int) {

	input := utils.ReadFileToString("2022/day13/input.txt")
	pairs := strings.Split(input, "\n\n")

	return solvePart1(pairs), solvePart2(pairs)
}

func solvePart1(pairs []string) int {
	res := 0

	for i, pairStr := range pairs {
		pair := strings.Split(pairStr, "\n")
		firstPart := parse(pair[0])
		secondPart := parse(pair[1])
		if compare(firstPart, secondPart) == -1 {
			res += i + 1
		}
	}

	return res
}

func solvePart2(pairs []string) int {
	packages := []string{"[[2]]", "[[6]]"}

	for _, pkg := range pairs {
		pair := strings.Split(pkg, "\n")
		first := pair[0]
		second := pair[1]
		packages = append(packages, first, second)
	}

	sort.Slice(packages, func(i, j int) bool {
		return compare(parse(packages[i]), parse(packages[j])) < 0
	})

	var twoIdx int
	var sixIdx int

	for i, curr := range packages {
		if curr == "[[2]]" {
			twoIdx = i + 1
		}
		if curr == "[[6]]" {
			sixIdx = i + 1
		}
	}

	return twoIdx * sixIdx
}

func compare(a any, b any) int {

	aIsList := isList(reflect.TypeOf(a))
	bIsList := isList(reflect.TypeOf(b))

	if !aIsList && !bIsList {
		return compareNums(a.(float64), b.(float64))
	}

	if aIsList && !bIsList {
		return compare(a, []any{b.(float64)})
	}

	if !aIsList && bIsList {
		return compare([]any{a.(float64)}, b)
	}

	aList := a.([]any)
	bList := b.([]any)

	lenA := len(aList)
	lenB := len(bList)
	smallerLen := utils.Min(lenA, lenB)

	for i := 0; i < smallerLen; i++ {
		res := compare(aList[i], bList[i])
		if res != 0 {
			return res
		}
	}
	return compareNums(float64(lenA), float64(lenB))
}

func compareNums(a float64, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func isList(t reflect.Type) bool {
	return strings.HasPrefix(t.String(), "[]")
}

func parse(input string) []any {
	var data []any
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		panic(err)
	}

	return data
}
