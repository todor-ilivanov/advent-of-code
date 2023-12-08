package day08

import (
	"advent/utils"
	"strings"
)

var MOVES = map[rune]int{
	'L': 0,
	'R': 1,
}

func moveGenerator(moves string) func() int {
	index := 0
	return func() int {
		nextMove := moves[index%len(moves)]
		index++
		return MOVES[rune(nextMove)]
	}
}

func Solve() (int, int) {
	input := utils.ReadFileToString("2023/day08/input.txt")
	rows := strings.Split(input, "\n")

	nodes := make(map[string][2]string)

	for _, row := range rows[2:] {
		tokens := strings.Split(row, " = ")

		node := tokens[0]
		neighboursStr := tokens[1]
		neighbours := strings.Split(neighboursStr[1:len(neighboursStr)-1], ", ")

		nodes[node] = [2]string{neighbours[0], neighbours[1]}
	}

	stepsSingle := countStepsSingle(rows[0], nodes)
	stepsMultiple := countStepsMuiltiple(rows[0], nodes)

	return stepsSingle, lcm(stepsMultiple[0], stepsMultiple[1], stepsMultiple...)
}

func countStepsSingle(moves string, nodes map[string][2]string) int {

	nextMove := moveGenerator(moves)
	current := "AAA"
	steps := 0

	for current != "ZZZ" {
		move := nextMove()
		next := nodes[current][move]
		current = next
		steps++
	}
	return steps
}

func countStepsMuiltiple(moves string, nodes map[string][2]string) []int {

	currentNodes := make([]string, 0)

	for k := range nodes {
		if k[2] == 'A' {
			currentNodes = append(currentNodes, k)
		}
	}

	allSteps := make([]int, len(currentNodes))

	for i, node := range currentNodes {

		current := node
		steps := 0
		nextMove := moveGenerator(moves)

		for current[2] != 'Z' {
			move := nextMove()
			next := nodes[current][move]
			current = next
			steps++
		}

		allSteps[i] = steps
	}

	return allSteps
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
