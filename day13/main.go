package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Machine struct {
	A     helper.Coord
	B     helper.Coord
	Prize helper.Coord
}

func readData(lines []string) []Machine {
	r := make([]Machine, 0)
	var tmpMachine Machine
	for _, line := range lines {
		if strings.Contains(line, "Button") {
			x, _ := strconv.Atoi(strings.TrimSuffix(strings.Split(strings.Split(line, " ")[2], "+")[1], ","))
			y, _ := strconv.Atoi(strings.Split(strings.Split(line, " ")[3], "+")[1])
			if strings.Contains(line, "Button A") {
				tmpMachine = Machine{A: helper.Coord{X: x, Y: y}}
			} else {
				tmpMachine.B = helper.Coord{X: x, Y: y}
			}
		} else if strings.Contains(line, "Prize") {
			px, _ := strconv.Atoi(strings.TrimSuffix(strings.Split(strings.Split(line, " ")[1], "=")[1], ","))
			py, _ := strconv.Atoi(strings.Split(strings.Split(line, " ")[2], "=")[1])
			tmpMachine.Prize = helper.Coord{X: px, Y: py}
			r = append(r, tmpMachine)
		} else {
			continue
		}
	}
	return r
}

// Based on https://old.reddit.com/r/adventofcode/comments/1hd4wda/2024_day_13_solutions/m27uwd9/
// Can't take any credit for this one!
func divMod(dividend, divisor int) (quotient, remainder int) {
	if divisor == 0 {
		panic("division by zero is not allowed")
	}
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

func minCost(ax, ay, bx, by, px, py int) int {
	b, brem := divMod(ay*px-ax*py, ay*bx-ax*by)
	a, arem := divMod(px-b*bx, ax)
	if arem > 0 || brem > 0 {
		return 0
	} else {
		return a*3 + b
	}
}

func partone(lines []string) (r int, err error) {
	machines := readData(lines)

	for _, machine := range machines {
		r += minCost(machine.A.X, machine.A.Y, machine.B.X, machine.B.Y, machine.Prize.X, machine.Prize.Y)
	}
	return r, nil
}

func parttwo(lines []string) (r int, err error) {
	machines := readData(lines)

	for _, machine := range machines {
		r += minCost(machine.A.X, machine.A.Y, machine.B.X, machine.B.Y, 10000000000000+machine.Prize.X, 10000000000000+machine.Prize.Y)
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
