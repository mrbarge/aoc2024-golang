package main

import (
	"fmt"
	"math"
	"os"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/mrbarge/aoc2024-golang/helper"
)

type Walk struct {
	pos  helper.Coord
	cost int
}

func readData(lines []string) (r [][]bool, start helper.Coord, end helper.Coord) {
	r = make([][]bool, 0)
	for y, line := range lines {
		row := make([]bool, 0)
		for x, v := range line {
			row = append(row, v == '#')
			if v == 'S' {
				start = helper.Coord{X: x, Y: y}
			} else if v == 'E' {
				end = helper.Coord{X: x, Y: y}
			}
		}
		r = append(r, row)
	}
	return r, start, end
}

func byCost(a, b interface{}) int {
	priorityA := a.(Walk).cost
	priorityB := b.(Walk).cost
	return priorityA - priorityB
}

func shortest(grid [][]bool, start helper.Coord, end helper.Coord) (map[helper.Coord]int, int) {
	priorityQueue := priorityqueue.NewWith(byCost)

	distance := make(map[helper.Coord]int)
	best := 0
	priorityQueue.Enqueue(Walk{
		pos:  start,
		cost: 0,
	})

	for !priorityQueue.Empty() {
		elem, _ := priorityQueue.Dequeue()
		walk := elem.(Walk)

		if _, ok := distance[walk.pos]; !ok {
			distance[walk.pos] = walk.cost
		} else {
			continue
		}

		if walk.pos == end {
			best = walk.cost
		}

		for _, n := range walk.pos.GetNeighbours(false) {
			if n.IsValid(len(grid[0]), len(grid)) && !grid[n.Y][n.X] {
				priorityQueue.Enqueue(Walk{
					cost: walk.cost + 1,
					pos:  n,
				})
			}
		}
	}

	return distance, best
}

func distance(i helper.Coord, j helper.Coord) int {
	return int(math.Abs(float64(i.X-j.X)) + math.Abs(float64(i.Y-j.Y)))
}

func partone(lines []string, allowedCheat int) (r int, err error) {
	grid, start, end := readData(lines)
	bestpathFromStart, score := shortest(grid, start, end)
	bestpathFromEnd, _ := shortest(grid, end, start)

	for y, row := range grid {
		for x, _ := range row {
			csc := helper.Coord{X: x, Y: y}
			_, ok := bestpathFromStart[csc]
			if grid[y][x] || !ok {
				// it's a wall, or it's not on our best path
				continue
			}
			// ranges for cheating
			cheaty1 := int(math.Max(float64(y-allowedCheat), 0))
			cheaty2 := int(math.Min(float64(y+allowedCheat), float64(len(grid)-1)))
			cheatx1 := int(math.Max(float64(x-allowedCheat), 0))
			cheatx2 := int(math.Min(float64(x+allowedCheat), float64(len(row)-1)))

			for iy := cheaty1; iy <= cheaty2; iy++ {
				for ix := cheatx1; ix <= cheatx2; ix++ {
					cse := helper.Coord{X: ix, Y: iy}
					if grid[iy][ix] == true {
						continue
					}
					if _, ok := bestpathFromEnd[cse]; !ok {
						continue
					}
					dist := distance(helper.Coord{X: x, Y: y}, helper.Coord{X: ix, Y: iy})
					if dist > allowedCheat {
						continue
					}

					cheatSave := bestpathFromStart[csc] + bestpathFromEnd[cse] + dist
					if cheatSave <= score-100 {
						r++
					}
				}
			}
		}
	}
	return r, nil
}

func parttwo(lines []string) (r int, err error) {
	return partone(lines, 20)
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partone(lines, 2)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = parttwo(lines)
	fmt.Printf("Part two: %v\n", ans)

}
