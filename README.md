# Advent of Code

Go repository for the solutions of AOC over the years.

## Boilerplate for each day

1. Paste input in `input.txt`

2. `run.go` - dayN's folder:

```go
package dayN

import (
	"fmt"
)

func Run() {
	fmt.Println("----Day N----")
	part1, part2 := SolvePart1(), SovlvePart2()
	fmt.Println(part1)
	fmt.Println(part2)
}

```

3. Implement the solutions in the day's package

4. `main.go` - in the project root, add a new case for the year/day:

```go
package main

import (
	...
	y22d01 "advent/2022/day01"
	y23d01 "advent/2023/day01"
)

func main() {

	...

	switch year {
	case "2022":
		switch day {
		case "01":
			y22d01.Run()
		}
		...
	case "2023":
		switch day {
		case "01":
			y23d01.Run()
		}
		...
	default:
		fmt.Println("Year or  not recognized")
	}
}
```

5. `go run main.go 2022 01`
