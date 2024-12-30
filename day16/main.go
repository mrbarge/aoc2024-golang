package main

import (
	"fmt"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/mrbarge/aoc2024-golang/helper"
	"math"
	"os"
	"strings"
)

type ProblemState struct {
	start helper.Coord
	end   helper.Coord
	grid  [][]bool
}

type Walk struct {
	pos  helper.Coord
	dir  helper.Direction
	cost int
	path []helper.Coord
}

type SimpleWalk struct {
	pos helper.Coord
	dir helper.Direction
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

func simulate(state ProblemState) int {
	priorityQueue := priorityqueue.NewWith(byCost)

	priorityQueue.Enqueue(Walk{
		pos:  state.start,
		dir:  helper.EAST,
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

		if walk.pos == state.end {
			// We arrived as quickly as we could
			return walk.cost
		}

		nextSteps := getNextSteps(walk, state, seen)
		for _, nextStep := range nextSteps {
			priorityQueue.Enqueue(nextStep)
		}

		seen[walk.pos] = walk.dir
	}
	return -1
}

func canVisit(w SimpleWalk, cost int, seen map[SimpleWalk]int) bool {
	if _, ok := seen[w]; ok {
		if seen[w] < cost {
			return false
		}
	}
	seen[w] = cost
	return true
}

func simulateTwo(state ProblemState) int {
	priorityQueue := priorityqueue.NewWith(byCost)

	priorityQueue.Enqueue(Walk{
		pos:  state.start,
		dir:  helper.EAST,
		cost: 0,
		path: []helper.Coord{state.start},
	})

	lowestCost := math.MaxInt

	seen := make(map[SimpleWalk]int)
	winning := make(map[helper.Coord]bool)
	for !priorityQueue.Empty() {
		elem, _ := priorityQueue.Dequeue()
		walk := elem.(Walk)

		if lowestCost < walk.cost {
			break
		}

		if walk.pos == state.end {
			lowestCost = walk.cost
			for _, v := range walk.path {
				winning[v] = true
			}
			continue
		}

		if !canVisit(SimpleWalk{pos: walk.pos, dir: walk.dir}, walk.cost, seen) {
			continue
		}

		forward := walk.pos.Move(walk.dir)
		newpath := make([]helper.Coord, len(walk.path))
		copy(newpath, walk.path)
		newpath = append(newpath, forward)
		if state.grid[forward.Y][forward.X] && canVisit(SimpleWalk{pos: forward, dir: walk.dir}, walk.cost+1, seen) {
			priorityQueue.Enqueue(Walk{
				pos:  forward,
				dir:  walk.dir,
				cost: walk.cost + 1,
				path: newpath,
			})
		}

		l := SimpleWalk{
			pos: walk.pos,
			dir: walk.dir.TurnAntiClockwise(),
		}
		r := SimpleWalk{
			pos: walk.pos,
			dir: walk.dir.TurnClockwise(),
		}
		for _, md := range []SimpleWalk{l, r} {
			if canVisit(md, walk.cost+1000, seen) {
				priorityQueue.Enqueue(Walk{
					pos:  md.pos,
					dir:  md.dir,
					cost: walk.cost + 1000,
					path: walk.path,
				})
			}
		}
	}

	return len(winning)
}

func pathToKey(a []helper.Coord) string {
	as := make([]string, 0)
	for _, v := range a {
		as = append(as, v.ToString())
	}
	return strings.Join(as, ":")
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
		newpath := make([]helper.Coord, 0)
		for _, v := range walk.path {
			newpath = append(newpath, v)
		}
		newpath = append(newpath, neighbour)
		nextSteps = append(nextSteps, Walk{
			pos:  neighbour,
			dir:  ndir,
			cost: walk.cost + turnCost,
			path: newpath,
		})
	}
	return nextSteps
}

func partone(lines []string) (r int, err error) {
	state := readData(lines)
	r = simulate(state)
	return r, nil
}

func parttwo(lines []string) (r int, err error) {
	state := readData(lines)
	r = simulateTwo(state)
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
