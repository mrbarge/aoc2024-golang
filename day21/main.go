// https://github.com/AllanTaylor314/AdventOfCode/blob/main/2024/21.py
package main

import (
	"fmt"
	"os"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type WorldState struct {
	numKeypad map[rune]helper.Coord
	numpad    [][]bool
	dirpad    [][]bool
	dirKeypad map[helper.Direction]helper.Coord

	numpadBot  helper.Coord
	dirpadBot  helper.Coord
	dirpadBot2 helper.Coord
}

func readData(lines []string) WorldState {
	w := WorldState{}
	w.numKeypad = map[rune]helper.Coord{
		'7': helper.Coord{},
		'8': helper.Coord{X: 1},
		'9': helper.Coord{X: 2},
		'4': helper.Coord{X: 0, Y: 1},
		'5': helper.Coord{X: 1, Y: 1},
		'6': helper.Coord{X: 2, Y: 1},
		'1': helper.Coord{Y: 2},
		'2': helper.Coord{X: 1, Y: 2},
		'3': helper.Coord{X: 2, Y: 2},
		'0': helper.Coord{X: 1, Y: 3},
		'A': helper.Coord{X: 2, Y: 3},
	}
	w.numpad = [][]bool{
		{true, true, true},
		{true, true, true},
		{true, true, true},
		{false, true, true},
	}
	w.dirpad = [][]bool{
		{false, true, true},
		{true, true, true},
	}
	w.dirKeypad = map[helper.Direction]helper.Coord{
		helper.NORTH:     helper.Coord{X: 1},
		helper.NORTHEAST: helper.Coord{X: 2},
		helper.EAST:      helper.Coord{X: 2, Y: 1},
		helper.SOUTH:     helper.Coord{X: 1, Y: 1},
		helper.WEST:      helper.Coord{Y: 1},
	}

	w.numpadBot = helper.Coord{X: 3, Y: 3}
	w.dirpadBot = helper.Coord{X: 2}
	w.dirpadBot2 = helper.Coord{X: 2}

	return w
}

func move(dest rune, w *WorldState) int {

	numpadDest := w.numpad[dest]

}

func partone(lines []string) (r int, err error) {
	worldstate := readData(lines)

	for _, line := range lines {
		for _, v := range line {
			moves := move(v, &worldstate)
			fmt.Printf("Moves for %v: %v\n", v, moves)
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
