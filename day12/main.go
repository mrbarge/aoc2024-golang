package main

import (
	"fmt"
	"os"

	"github.com/mrbarge/aoc2024-golang/helper"
)

func partone(lines []string) (r int, err error) {
	return 0, nil
}

func parttwo(lines []string) (r int, err error) {
	return 0, nil
}

func main() {
	fh, _ := os.Open("test.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partone(lines)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = parttwo(lines)
	fmt.Printf("Part two: %v\n", ans)

}
