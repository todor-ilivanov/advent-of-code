#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 year day"
    exit 1
fi

year=$1
day=$2
dirName="${year}/day${day}"

mkdir -p $dirName

touch $dirName/input.txt

cat << EOF > "${dirName}/sol.go"
package day${day}

import (
    "advent/utils"
    "fmt"
    "strings"
)

func Solve() (int, int) {
    input := utils.ReadFileToString("${year}/day${day}/input.txt")
    rows := strings.Split(input, "\n")

    fmt.Println(rows)
    return 0, 0
}
EOF

cat << EOF > "${dirName}/run.go"
package day${day}

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println("----Day ${day}----")

	start := time.Now()
	part1, part2 := Solve()
	fmt.Println(part1)
	fmt.Println(part2)

	fmt.Println("Time elapsed:", time.Since(start))
}
EOF