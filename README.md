# Advent of Code

Go repository for the solutions of AOC over the years.

## Boilerplate for each day

1. Run `newday.sh 2023 01` - creates a new dir and generates scaffold (`run.go`, `sol.go` and `input.txt`)

2. Input needs to be pasted manually

3. `main.go` - in the project root, add a new case for the year/day:

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
}
```

4. `go run main.go 2023 01`
