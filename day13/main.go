package main

import (
	"fmt"
	"math"
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

func walk(pos helper.Coord, aturn int, bturn int, tokens int, m Machine) int {
	if pos == m.Prize {
		return tokens
	}
	apos := helper.Coord{X: pos.X + m.A.X, Y: pos.Y + m.A.Y}
	bpos := helper.Coord{X: pos.X + m.B.X, Y: pos.Y + m.B.Y}

	var w1, w2 int = math.MaxInt, math.MaxInt
	if aturn < 100 {
		w1 = walk(apos, aturn+1, bturn, tokens+3, m)
	}
	if bturn < 100 {
		w2 = walk(bpos, aturn, bturn+1, tokens+1, m)
	}
	if w1 < w2 {
		return w1
	} else {
		return w2
	}
}

func partone(lines []string) (r int, err error) {
	machines := readData(lines)

	for i, machine := range machines {
		tokens := walk(helper.Coord{}, 0, 0, 0, machine)
		if tokens == math.MaxInt {
			fmt.Printf("Couldn't solve machine %v\n", i)
		} else {
			fmt.Printf("Machine %v tokens %v\n", i, tokens)
		}
	}
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
