# Advent of code 2022

## Boilerplate for each day

In day's folder .go file:
```go
package dayN

import (
	"fmt"
	"strings"

	"advent/utils"
)

func Solve() (int, int) {

	input := utils.ReadFileToString("dayN/input.txt")
	rows := strings.Split(input, "\n")

	fmt.Println(rows)

	return 0, 0
}
```

In main.go:
```go
package main

import (
	"fmt"

	"advent/dayN"
)

func main() {
	solveDayN()
}

func solveDayN() {
	fmt.Println("----Day N----")
	part1 := dayN.SolvePart1()
	fmt.Println(part1)
	part2 := dayN.SolvePart2(10)
	fmt.Println(part2)
}

```
