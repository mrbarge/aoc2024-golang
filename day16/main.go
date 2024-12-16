package main

import (
	"fmt"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/mrbarge/aoc2024-golang/helper"
	"os"
)

type ProblemState struct {
	start helper.Coord
	end   helper.Coord
	grid  [][]bool
}

type Walk struct {
	pos  helper.Coord
	dir  helper.Direction
	path map[helper.Coord]int
	cost int
}

// Comparator function
func byCost(a, b interface{}) int {
	priorityA := a.(Walk).cost
	priorityB := b.(Walk).cost
	return priorityA - priorityB
}

func readData(lines []string) ProblemState {
	state := ProblemState{
		grid: make([][]bool, 0),
	}
	for y, line := range lines {
		row := make([]bool, 0)
		for x, v := range line {
			if v == '#' {
				row = append(row, false)
			} else if v == 'E' {
				row = append(row, true)
				state.end = helper.Coord{X: x, Y: y}
			} else if v == 'S' {
				row = append(row, true)
				state.start = helper.Coord{X: x, Y: y}
			} else {
				row = append(row, true)
			}
		}
		state.grid = append(state.grid, row)
	}
	return state
}

func isTurn(facing helper.Direction, walking helper.Direction) bool {
	return facing != walking
}

func simulate(state ProblemState) (int, map[helper.Coord]int) {
	priorityQueue := priorityqueue.NewWith(byCost)

	priorityQueue.Enqueue(Walk{
		pos:  state.start,
		dir:  helper.EAST,
		path: make(map[helper.Coord]int),
		cost: 0,
	})

	seen := make(map[helper.Coord]helper.Direction)

	for !priorityQueue.Empty() {
		elem, _ := priorityQueue.Dequeue()
		walk := elem.(Walk)

		if _, ok := seen[walk.pos]; ok {
			if seen[walk.pos] == walk.dir {
				// we have walked here, in the same direction, before
				// so this is a loop
				continue
			}
		}

		// Add to the path
		walk.path[walk.pos] = walk.cost

		if walk.pos == state.end {
			// We arrived as quickly as we could
			return walk.cost, walk.path
		}

		nextSteps := getNextSteps(walk, state, seen)
		for _, nextStep := range nextSteps {
			priorityQueue.Enqueue(nextStep)
		}

		seen[walk.pos] = walk.dir
	}
	return -1, nil
}

func simulateTwo(state ProblemState) (int, int) {
	priorityQueue := priorityqueue.NewWith(byCost)

	priorityQueue.Enqueue(Walk{
		pos:  state.start,
		dir:  helper.EAST,
		path: make(map[helper.Coord]int),
		cost: 0,
	})

	return 0, 0
}
func getNextSteps(walk Walk, state ProblemState, seen map[helper.Coord]helper.Direction) []Walk {

	nextSteps := make([]Walk, 0)
	neighbours := walk.pos.GetNeighboursAsMap(false)
	for ndir, neighbour := range neighbours {
		if !neighbour.IsValid(len(state.grid[0]), len(state.grid)) {
			// off the map somehow
			continue
		} else if !state.grid[neighbour.Y][neighbour.X] {
			// a wall
			continue
		} else if ndir.Opposite() == walk.dir {
			// We'd be going backwards..
			continue
		} else if _, ok := seen[neighbour]; ok {
			if seen[neighbour] == ndir {
				// a repeat
				continue
			}
		}
		turning := isTurn(walk.dir, ndir)
		turnCost := 1
		if turning {
			turnCost += 1000
		}
		nextSteps = append(nextSteps, Walk{
			pos:  neighbour,
			dir:  ndir,
			path: copyPath(walk.path),
			cost: walk.cost + turnCost,
		})
	}
	return nextSteps
}

func copyPath(path map[helper.Coord]int) (r map[helper.Coord]int) {
	r = make(map[helper.Coord]int)
	for i, v := range path {
		r[i] = v
	}
	return r
}

func partone(lines []string) (r int, err error) {
	state := readData(lines)
	r, _ = simulate(state)
	return r, nil
}

func parttwo(lines []string) (r int, err error) {
	state := readData(lines)
	_, tiles := simulateTwo(state)
	return tiles, nil
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
