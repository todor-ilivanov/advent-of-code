package day11

import (
	"regexp"
	"sort"
	"strings"

	"advent/utils"
)

type operation struct {
	sign  string
	param int
}

type test struct {
	divisibleBy   int
	monkeyIfTrue  int
	monkeyIfFalse int
}

type monkey struct {
	num   int
	items []int
	op    operation
	test  test
}

var m int = 1

func Solve(rounds int, divisor int) int {

	input := utils.ReadFileToString("2022/day11/input.txt")
	monkeysRaw := strings.Split(input, "Monkey")

	monkeys, monkeyInspections := loadMonkeys(monkeysRaw)

	for i := 0; i < rounds; i++ {
		processItems(monkeys, monkeyInspections, divisor)
	}

	var inspections []int
	for _, v := range monkeyInspections {
		inspections = append(inspections, v)
	}

	sort.Ints(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func processItems(monkeys []*monkey, monkeyInspections map[int]int, divisor int) {
	for _, monkey := range monkeys {

		for _, val := range monkey.items {
			currValue := calcNewValue(val, monkey.op)
			nextValue := currValue / divisor

			if nextValue%monkey.test.divisibleBy == 0 {
				nextMonkey := findMonkeyByNum(monkey.test.monkeyIfTrue, monkeys)
				nextMonkey.items = append(nextMonkey.items, nextValue)
			} else {
				nextMonkey := findMonkeyByNum(monkey.test.monkeyIfFalse, monkeys)
				nextMonkey.items = append(nextMonkey.items, nextValue)
			}

			monkey.items = monkey.items[1:]

			monkeyInspections[monkey.num]++
		}
	}
}

func calcNewValue(val int, op operation) int {
	param := op.param
	if param == -1 {
		param = val
	}

	switch op.sign {
	case "*":
		return (val * param) % m
	case "/":
		return val / param
	case "+":
		return (val + param) % m
	case "-":
		return val - param
	default:
		return val
	}
}

func findMonkeyByNum(num int, monkeys []*monkey) *monkey {
	for _, monkey := range monkeys {
		if monkey.num == num {
			return monkey
		}
	}
	return nil
}

func loadMonkeys(monkeysRaw []string) ([]*monkey, map[int]int) {
	monkeys := make([]*monkey, 0)
	monkeyInspections := make(map[int]int)

	reNum := regexp.MustCompile("[0-9]+")
	reSign := regexp.MustCompile("[+-/*]")

	for _, monkeyRaw := range monkeysRaw[1:] {

		monkey := monkey{}
		monkeyRawCleaned := strings.TrimSuffix(monkeyRaw, "\n")
		monkeyRawCleaned = strings.TrimSpace(monkeyRawCleaned)

		for _, line := range strings.Split(monkeyRawCleaned, "\n") {

			lineCleaned := strings.TrimSpace(line)
			numsAsStr := reNum.FindAllString(line, -1)

			if strings.HasPrefix(lineCleaned, "Starting items") {
				for _, numStr := range numsAsStr {
					num := utils.StringToInt(numStr)
					monkey.items = append(monkey.items, num)
				}
			} else if strings.HasPrefix(lineCleaned, "Operation") {
				monkey.op.sign = reSign.FindAllString(line, -1)[0]

				if len(numsAsStr) > 0 {
					monkey.op.param = utils.StringToInt(numsAsStr[0])
				} else {
					monkey.op.param = -1 // denotes "old * old"
				}
			} else if strings.HasPrefix(lineCleaned, "Test") {
				monkey.test.divisibleBy = utils.StringToInt(numsAsStr[0])
			} else if strings.HasPrefix(lineCleaned, "If true") {
				monkey.test.monkeyIfTrue = utils.StringToInt(numsAsStr[0])
			} else if strings.HasPrefix(lineCleaned, "If false") {
				monkey.test.monkeyIfFalse = utils.StringToInt(numsAsStr[0])
			} else {
				monkey.num = utils.StringToInt(numsAsStr[0])
			}
		}

		m *= monkey.test.divisibleBy
		monkeyInspections[monkey.num] = 0
		monkeys = append(monkeys, &monkey)
	}

	return monkeys, monkeyInspections
}
