package day05

import (
	"advent/utils"
	"math"
	"strings"
	"unicode"
)

type mapping struct {
	converters []*converter
	name       string
}

func newMapping(name string) *mapping {
	converters := make([]*converter, 0)
	mapping := mapping{name: name, converters: converters}
	return &mapping
}

func (m mapping) mapNum(num int) int {
	originalNum := num
	for _, c := range m.converters {
		converted := c.convert(num)
		if converted != originalNum {
			return converted
		}
	}
	return num
}

type converter struct {
	offset int
	min    int
	max    int
}

func newConverter(destination, source, rangeMax int) *converter {
	offset := destination - source
	max := source + rangeMax - 1
	converter := converter{min: source, max: max, offset: offset}
	return &converter
}

func (c converter) convert(num int) int {
	if num >= c.min && num <= c.max {
		return num + c.offset
	}
	return num
}

func Solve() (int, int) {
	input := utils.ReadFileToString("2023/day05/input.txt")
	rows := strings.Split(input, "\n")

	seeds := parseSeeds(rows[0])
	mappings := parseMappings(rows[1:])

	return minLocationForSeeds(seeds, mappings), 0 // bruteForcePart2(seeds, mappings)
}

func parseSeeds(raw string) []int {
	tokens := strings.Split(raw, ": ")
	return convertRowToIntSlice(tokens[1])
}

func bruteForcePart2(seeds []int, mappings []*mapping) int {

	minLoc := math.MaxInt
	i, j := 0, 1

	for j < len(seeds) {
		start := seeds[i]
		end := start + seeds[j]

		for start < end {
			res := start
			for _, m := range mappings {
				res = m.mapNum(res)
			}
			minLoc = utils.Min(minLoc, res)
			start += 1
		}

		i += 2
		j += 2
	}

	return minLoc
}

func parseMappings(rows []string) []*mapping {

	mappings := make([]*mapping, 0)
	var currentMapping *mapping

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		firstChar := rune(row[0])

		if unicode.IsLetter(firstChar) {
			currentMapping = newMapping(row)
			mappings = append(mappings, currentMapping)
		} else if unicode.IsDigit(firstChar) {
			nums := convertRowToIntSlice(row)
			converter := newConverter(nums[0], nums[1], nums[2])
			currentMapping.converters = append(currentMapping.converters, converter)
		}
	}

	return mappings
}

func minLocationForSeeds(seeds []int, mappings []*mapping) int {

	minLoc := math.MaxInt

	for _, seed := range seeds {
		res := seed
		for _, m := range mappings {
			res = m.mapNum(res)
		}
		minLoc = utils.Min(minLoc, res)
	}

	return minLoc
}

func convertRowToIntSlice(row string) []int {
	slice := make([]int, 0)

	for _, numStr := range strings.Split(row, " ") {
		num := utils.StringToInt(numStr)
		slice = append(slice, num)
	}

	return slice
}
