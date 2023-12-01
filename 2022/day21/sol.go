package day21

import (
	"regexp"
	"strings"

	"advent/utils"
)

type Monkey struct {
	name     string
	children [2]*Monkey
	op       string
	val      int
}

func Solve() (int, int) {

	input := utils.ReadFileToString("2022/day21/input.txt")
	rows := strings.Split(input, "\n")

	monkeys := populateMonkeysMap(rows)

	part1 := calcValue(monkeys["root"])

	return part1, 0
}

func populateMonkeysMap(rows []string) map[string]*Monkey {
	monkeys := make(map[string]*Monkey)

	for _, row := range rows {

		tokens := strings.Split(row, " ")
		name := utils.Trim(tokens[0])

		monkey := getMonkeyOrCreateNew(monkeys, name)

		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(row, -1)
		if len(nums) > 0 {
			monkey.val = utils.StringToInt(nums[0])
			monkeys[name] = monkey
			continue
		}

		leftName := tokens[1]
		op := tokens[2]
		rightName := tokens[3]

		if contains(monkeys, leftName) {
			monkey.children[0] = monkeys[leftName]
		} else {
			leftMonkey := Monkey{leftName, [2]*Monkey{}, "", 0}
			monkey.children[0] = &leftMonkey
			monkeys[leftName] = &leftMonkey
		}

		if contains(monkeys, rightName) {
			monkey.children[1] = monkeys[rightName]
		} else {
			rightMonkey := Monkey{rightName, [2]*Monkey{}, "", 0}
			monkey.children[1] = &rightMonkey
			monkeys[rightName] = &rightMonkey
		}

		monkey.op = op
		monkeys[name] = monkey
	}

	return monkeys
}

func calcValue(monkey *Monkey) int {

	if monkey.val != 0 {
		return monkey.val
	}

	monkey.children[0].val = calcValue(monkey.children[0])
	monkey.children[1].val = calcValue(monkey.children[1])

	return calc(monkey.children[0].val, monkey.children[1].val, monkey.op)
}

func calc(val1, val2 int, op string) int {
	switch op {
	case "*":
		return val1 * val2
	case "/":
		return val1 / val2
	case "+":
		return val1 + val2
	case "-":
		return val1 - val2
	default:
		panic("Unsupported operation.")
	}
}

func getMonkeyOrCreateNew(monkeys map[string]*Monkey, name string) *Monkey {
	if contains(monkeys, name) {
		return monkeys[name]
	} else {
		return &Monkey{name, [2]*Monkey{}, "", 0}
	}
}

func contains(monkeys map[string]*Monkey, name string) bool {
	if _, ok := monkeys[name]; ok {
		return ok
	}
	return false
}
