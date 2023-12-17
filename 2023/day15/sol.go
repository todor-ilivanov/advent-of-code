package day15

import (
	"advent/utils"
	"fmt"
	"strings"
	"unicode"
)

type lens struct {
	key string
	val int
}

type listNode struct {
	val  *lens
	next *listNode
}

func Solve() (int, int) {
	input := utils.ReadFileToString("2023/day15/input.txt")
	commands := strings.Split(input, ",")

	hashSum := calculateHashSum(commands)

	hashMap := make([]*listNode, 256)

	for _, c := range commands {

		label, operation, value := parseCommand(c)

		if operation == '=' {
			addLens(hashMap, label, value)
		} else {
			removeLens(hashMap, label)
		}
	}

	printHashMap(hashMap)

	return hashSum, calculateFocusingPower(hashMap)
}

func calculateHashSum(commands []string) int {
	sum := 0
	for _, c := range commands {
		sum += hash(c)
	}
	return sum
}

func hash(s string) int {
	value := 0
	for _, ch := range s {
		value += int(ch)
		value *= 17
		value %= 256
	}
	return value
}

func parseCommand(s string) (string, rune, int) {
	var label []string
	var operation rune
	var value int

	for _, ch := range s {
		if !unicode.IsLetter(ch) {
			operation = ch
			break
		}
		label = append(label, string(ch))
	}

	if operation == '=' {
		digitStr := string(s[len(s)-1])
		value = utils.StringToInt(digitStr)
	}

	return strings.Join(label, ""), operation, value
}

func addLens(hashMap []*listNode, label string, value int) {
	containsLabel := false
	index := hash(label)
	current := hashMap[index]
	var prev *listNode

	for current != nil {
		lens := current.val
		if lens.key == label {
			lens.val = value
			containsLabel = true
			break
		}
		prev = current
		current = current.next
	}

	if !containsLabel {
		if prev != nil {
			prev.next = &listNode{&lens{label, value}, nil}
		} else {
			hashMap[index] = &listNode{&lens{label, value}, nil}
		}
	}
}

func removeLens(hashMap []*listNode, label string) {

	index := hash(label)

	if hashMap[index] == nil {
		return
	}

	if hashMap[index].val.key == label {
		hashMap[index] = hashMap[index].next
		return
	}

	current := hashMap[index]
	prev := hashMap[index]

	for current != nil {
		lens := current.val
		if lens.key == label {
			prev.next = current.next
			break
		}
		prev = current
		current = current.next
	}
}

func calculateFocusingPower(hashMap []*listNode) int {

	score := 0

	for i := range hashMap {
		if hashMap[i] == nil {
			continue
		}

		current := hashMap[i]
		listIdx := 1
		for current != nil {
			lens := *current.val
			score += (i + 1) * listIdx * lens.val
			listIdx++
			current = current.next
		}
	}

	return score
}

func printHashMap(hashMap []*listNode) {
	for i := range hashMap {
		if hashMap[i] == nil {
			continue
		}
		fmt.Println("----- box ", i)
		current := hashMap[i]
		for current != nil {
			fmt.Println(current.val.key, current.val.val)
			current = current.next
		}
	}
}
