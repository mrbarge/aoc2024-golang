package main

import (
	"fmt"
	"github.com/mrbarge/aoc2024-golang/helper"
	"math"
	"os"
	"sort"
)

func getSimilarity(x []int, y []int) (r map[int]int) {
	r = make(map[int]int)
	for _, i := range x {
		if _, ok := r[i]; !ok {
			r[i] = 0
		} else {
			continue
		}
		for _, j := range y {
			if i == j {
				r[i] += 1
			}
		}
	}
	return r
}

func getSortedSides(data [][]int) ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)
	for _, line := range data {
		left = append(left, line[0])
		right = append(right, line[1])
	}
	sort.Ints(left)
	sort.Ints(right)

	return left, right
}

func problem(lines [][]int, partTwo bool) (r float64, err error) {

	left, right := getSortedSides(lines)

	if !partTwo {
		for i, v := range left {
			r += math.Abs(float64(v) - float64(right[i]))
		}
		return r, nil
	} else {
		similarity := getSimilarity(left, right)
		for _, v := range left {
			r += float64(v * similarity[v])
		}
		return r, nil
	}
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLinesAsCSVIntArray(fh, "   ")
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, false)
	fmt.Printf("Part one: %f\n", ans)

	ans, err = problem(lines, true)
	fmt.Printf("Part two: %f\n", ans)

}
