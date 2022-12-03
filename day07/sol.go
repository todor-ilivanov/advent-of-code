package day07

import (
	"sort"
	"strings"

	"advent/utils"
)

type dir struct {
	parent     *dir
	directSize int
	name       string
	children   []*dir
}

const totalAvailableSpace int = 70000000
const spaceNeededForUpdate int = 30000000
const maxSize int = 100000

var totalSizes map[*dir]int = make(map[*dir]int)

func Solve() (int, int) {

	input := utils.ReadFileToString("day07/input.txt")

	rows := strings.Split(input, "\n")

	var root dir = dir{nil, 0, "root", make([]*dir, 0)}

	parseInputAndPopulateTree(rows, &root)

	totalSpaceTaken := getTotalSizeForDir(root)

	spaceNeeded := spaceNeededForUpdate - (totalAvailableSpace - totalSpaceTaken)

	var resultPart1 int
	var resultPart2 []int

	for _, v := range totalSizes {
		if v <= maxSize {
			resultPart1 += v
		}
		if v >= spaceNeeded {
			resultPart2 = append(resultPart2, v)
		}
	}

	sort.Ints(resultPart2)

	return resultPart1, resultPart2[0]
}

func parseInputAndPopulateTree(rows []string, root *dir) {
	current := root
	for i := 1; i < len(rows)-1; i++ {
		row := rows[i]
		command := strings.Split(row, " ")[1]
		if command == "ls" {
			var lsRows []string = getListedItems(rows[i+1:])
			i += len(lsRows)
			for _, lsRow := range lsRows {
				elements := strings.Split(lsRow, " ")
				dirOrSize := string(elements[0])
				if dirOrSize == "dir" {
					dirName := string(elements[1])
					child := dir{current, 0, dirName, make([]*dir, 0)}
					current.children = append(current.children, &child)
				} else {
					current.directSize += utils.StringToInt(dirOrSize)
				}
			}
		}

		if command == "cd" {
			param := strings.Split(row, " ")[2]
			if param == ".." {
				current = current.parent
				continue
			}

			for _, child := range current.children {
				if child.name == param {
					current = child
					break
				}
			}
		}
	}
}

func getListedItems(rows []string) []string {
	var lsRows []string
	for _, res := range rows {
		if isCommand(res) {
			return lsRows
		}
		lsRows = append(lsRows, res)
	}
	return lsRows
}

func isCommand(row string) bool {
	return string(row[0]) == "$"
}

func getTotalSizeForDir(directory dir) int {
	if len(directory.children) == 0 {
		totalSizes[&directory] = directory.directSize
		return directory.directSize
	}

	if val, ok := totalSizes[&directory]; ok {
		return val
	}

	sumChildSizes := 0
	for _, child := range directory.children {
		sumChildSizes += getTotalSizeForDir(*child)
	}
	totalSizes[&directory] = directory.directSize + sumChildSizes
	return totalSizes[&directory]
}
