package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Sum struct {
	answer  int
	numbers []int
}

func readData(lines []string) []Sum {
	r := make([]Sum, 0)
	for _, line := range lines {
		left, _ := strconv.Atoi(strings.Split(line, ": ")[0])
		rrow := make([]int, 0)
		right := strings.Split(strings.Split(line, ": ")[1], " ")
		for _, v := range right {
			iv, _ := strconv.Atoi(v)
			rrow = append(rrow, iv)
		}
		s := Sum{answer: left, numbers: rrow}
		r = append(r, s)
	}
	return r
}

func (s Sum) IsValid(partTwo bool) bool {
	if !partTwo {
		return applyOperator(s.numbers, 0, s.answer, 0)
	} else {
		return applyOperatorTwo(s.numbers, 1, s.answer, s.numbers[0])
	}
}

func applyOperator(numbers []int, idx int, expected int, total int) bool {
	if idx == len(numbers) {
		return expected == total
	}
	mult := applyOperator(numbers, idx+1, expected, total*numbers[idx])
	add := applyOperator(numbers, idx+1, expected, total+numbers[idx])
	return mult || add
}

func applyOperatorTwo(numbers []int, idx int, expected int, total int) bool {
	if idx == len(numbers) {
		return expected == total
	}
	mult := applyOperatorTwo(numbers, idx+1, expected, total*numbers[idx])
	add := applyOperatorTwo(numbers, idx+1, expected, total+numbers[idx])
	sconcatTotal := strconv.Itoa(total) + strconv.Itoa(numbers[idx])
	concatTotal, _ := strconv.Atoi(sconcatTotal)
	concat := applyOperatorTwo(numbers, idx+1, expected, concatTotal)
	return mult || add || concat
}

func partone(lines []string) (r int, err error) {
	sums := readData(lines)
	for _, sum := range sums {
		if sum.IsValid(false) {
			r += sum.answer
		}
	}
	return r, nil

}

func parttwo(lines []string) (r int, err error) {
	sums := readData(lines)
	for _, sum := range sums {
		if sum.IsValid(true) {
			r += sum.answer
		}
	}
	return r, nil
}

func main() {
	fh, _ := os.Open("input.txt")
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
