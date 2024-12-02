package main

import (
	"fmt"
	"math"
	"os"

	"github.com/mrbarge/aoc2024-golang/helper"
)

func inRange(i int, j int) bool {
	diff := math.Abs(float64(i - j))
	return diff >= 1 && diff <= 3
}

func lineSafe(line []int) bool {
	dirCount := make(map[bool]int)
	for i := 1; i < len(line); i++ {
		if !inRange(line[i-1], line[i]) {
			return false
		}
		dirCount[line[i] > line[i-1]]++
	}
	return !(dirCount[true] > 0 && dirCount[false] > 0)
}

func partone(lines [][]int) (r int, err error) {

	numSafe := 0
	for _, line := range lines {
		if lineSafe(line) {
			numSafe++
		}
	}
	return numSafe, nil
}

func permissive_line_safe(line []int) bool {
	for i := 0; i < len(line); i++ {
		// try all permutations of the line with one element removed
		newLine := append([]int{}, line[:i]...)
		newLine = append(newLine, line[i+1:]...)
		if lineSafe(newLine) {
			return true
		}
	}
	return false
}

func parttwo(lines [][]int) (r int, err error) {
	numSafe := 0
	for i, line := range lines {
		if permissive_line_safe(line) {
			fmt.Printf("Line %v safe: %v\n", i, line)
			numSafe++
		}
	}
	return numSafe, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLinesAsCSVIntArray(fh, " ")
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partone(lines)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = parttwo(lines)
	fmt.Printf("Part two: %v\n", ans)

}
