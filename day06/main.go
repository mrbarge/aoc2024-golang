package main

import (
	"fmt"
	"os"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Player struct {
	direction helper.Direction
	position  helper.Coord
}

func (p *Player) move(grid [][]bool) bool {

	nextPos := p.position.Move(p.direction)
	if !nextPos.IsValid(len(grid[0]), len(grid)) {
		return false
	}
	if grid[nextPos.Y][nextPos.X] {
		// hit an object, turn
		p.direction = p.direction.TurnClockwise()
	} else {
		p.position = nextPos
	}
	return true
}

func readData(lines []string) (r [][]bool, start helper.Coord) {
	r = make([][]bool, len(lines))
	for y, line := range lines {
		row := make([]bool, len(line))
		for x, v := range line {
			if v == '#' {
				row[x] = true
			} else if v == '^' {
				start.X = x
				start.Y = y
			}
		}
		r[y] = row
	}
	return r, start
}

func partone(lines []string) (r int, err error) {
	grid, startcoord := readData(lines)
	player := Player{
		direction: helper.NORTH,
		position:  startcoord,
	}
	seen := make(map[helper.Coord]bool)
	for {
		seen[player.position] = true
		if !player.move(grid) {
			break
		}
	}
	return len(seen), nil
}

func parttwo(lines []string) (r int, err error) {
	return 0, nil
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
