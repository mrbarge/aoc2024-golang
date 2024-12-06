package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/mrbarge/aoc2024-golang/helper"
)

func input_to_runes(input []string) (r [][]rune) {
	r = make([][]rune, len(input))
	for i, line := range input {
		r2 := make([]rune, len(line))
		for j, v := range line {
			r2[j] = rune(v)
		}
		r[i] = r2
	}
	return r
}

func isMas(c helper.Coord, data [][]rune) bool {

	move_pairs := [][]helper.Direction{
		{helper.NORTHWEST, helper.SOUTHEAST},
		{helper.NORTHEAST, helper.SOUTHWEST},
	}

	valid := true
	for _, pair := range move_pairs {
		neighbours := make([]rune, 0)
		for _, dir := range pair {
			neighbour := c.Move(dir)
			if neighbour.IsValid(len(data[0]), len(data)) {
				neighbours = append(neighbours, data[neighbour.Y][neighbour.X])
			}
		}
		if !(slices.Contains(neighbours, 'S') &&
			slices.Contains(neighbours, 'M')) {
			valid = false
		}
	}
	return valid
}

func countXmas(c helper.Coord, data [][]rune) (r int) {
	xmas := []rune{'M', 'A', 'S'}

	move_funcs := []helper.Direction{
		helper.NORTH, helper.SOUTH, helper.EAST, helper.WEST, helper.NORTHWEST, helper.NORTHEAST, helper.SOUTHWEST, helper.SOUTHEAST,
	}

	// Check left
	for _, fn := range move_funcs {
		moveCoord := helper.Coord{X: c.X, Y: c.Y}
		moveCoord = moveCoord.Move(fn)
		i := 0
		for ; moveCoord.IsValid(len(data[0]), len(data)) && i < 3; i += 1 {
			if data[moveCoord.Y][moveCoord.X] != xmas[i] {
				break
			}
			moveCoord = moveCoord.Move(fn)
		}
		if i == 3 {
			r++
		}
	}
	return r
}

func partone(lines []string) (r int, err error) {
	data := input_to_runes(lines)
	for y, row := range lines {
		for x, v := range row {
			if v != 'X' {
				continue
			}
			c := helper.Coord{X: x, Y: y}
			r += countXmas(c, data)
		}
	}
	return r, nil
}

func parttwo(lines []string) (r int, err error) {
	data := input_to_runes(lines)
	for y, row := range lines {
		for x, v := range row {
			if v != 'A' {
				continue
			}
			c := helper.Coord{X: x, Y: y}
			if isMas(c, data) {
				r += 1
			}
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
