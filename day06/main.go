package main

import (
	"fmt"
	"os"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Player struct {
	direction helper.Direction
	position  helper.Coord
	visited   map[string]bool
}

func (p *Player) move(grid [][]bool) (validMove bool, inLoop bool) {
	visitKey := fmt.Sprintf("%v,%v,%v", p.position.X, p.position.Y, p.direction)
	if _, ok := p.visited[visitKey]; !ok {
		p.visited[visitKey] = true
	} else {
		// we're at the same position we've been, moving in the same direction
		// so this has to be a loop...... right...?
		return true, true
	}

	nextPos := p.position.Move(p.direction)
	if !nextPos.IsValid(len(grid[0]), len(grid)) {
		return false, false
	}
	if grid[nextPos.Y][nextPos.X] {
		// hit an object, turn
		p.direction = p.direction.TurnClockwise()
	} else {
		p.position = nextPos
	}
	return true, false
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
	visited, _ := simulate(grid, startcoord)
	return visited, nil
}

func simulate(grid [][]bool, startcoord helper.Coord) (visited int, isloop bool) {
	player := Player{
		direction: helper.NORTH,
		position:  startcoord,
		visited:   make(map[string]bool),
	}
	seen := make(map[helper.Coord]bool)
	for {
		seen[player.position] = true
		if valid, looped := player.move(grid); !valid || looped {
			if looped {
				isloop = true
			}
			break
		}
	}
	return len(seen), isloop
}

func parttwo(lines []string) (r int, err error) {
	grid, startcoord := readData(lines)

	for y, row := range grid {
		for x, v := range row {
			if v || (y == startcoord.Y && x == startcoord.X) {
				// can't mess with this grid sapce
				continue
			}
			newgrid := copyGrid(grid)
			newgrid[y][x] = true
			_, loop := simulate(newgrid, startcoord)
			if loop {
				r += 1
			}

		}
	}
	return r, nil
}

func copyGrid(grid [][]bool) [][]bool {
	retGrid := make([][]bool, len(grid))
	for i := range grid {
		retGrid[i] = make([]bool, len(grid[i]))
		copy(retGrid[i], grid[i])
	}
	return retGrid
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
