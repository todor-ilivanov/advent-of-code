package day05

import (
	"strings"

	"advent/utils"
)

// todo remove hardcoding..
const NumStacks int = 9

func InitStacksFromFile(filename string) [NumStacks][]string {
	strStacks := utils.ReadFileToString(filename)
	rows := strings.Split(strStacks, "\n")

	var stacks [NumStacks][]string

	for i, row := range rows {
		stacks[i] = strings.Split(row, "")
	}

	return stacks
}
