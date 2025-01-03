package main

import (
	"fmt"
	"os"

	"github.com/albertorestifo/dijkstra"
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

	numpadGraph dijkstra.Graph
	dirpadGraph dijkstra.Graph
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

	w.numpadBot = helper.Coord{X: 2, Y: 3}
	w.dirpadBot = helper.Coord{X: 2}
	w.dirpadBot2 = helper.Coord{X: 2}
	w.numpadGraph = padGraph(w.numpad)
	w.dirpadGraph = padGraph(w.dirpad)
	return w
}

func pathToDirections(path []string) []helper.Direction {
	r := make([]helper.Direction, 0)
	for x := 0; x < len(path)-1; x++ {
		c1 := helper.ParseCoord(path[x])
		c2 := helper.ParseCoord(path[x+1])
		r = append(r, c1.Direction(c2))
	}
	return r
}

func move(dest rune, w *WorldState) int {

	// First we need to figure out the shortest path for the numpad robot
	fmt.Printf("Moving from %v to %v\n", w.numpadBot.ToString(), w.numKeypad[dest].ToString())
	numpadPath, _, _ := w.numpadGraph.Path(w.numpadBot.ToString(), w.numKeypad[dest].ToString())
	fmt.Printf("%v\n", pathToDirections(numpadPath))
	return 0
}

func padGraph(grid [][]bool) dijkstra.Graph {
	g := dijkstra.Graph{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if !grid[y][x] {
				continue
			}
			c := helper.Coord{X: x, Y: y}
			nodeVal := fmt.Sprintf("%v,%v", x, y)
			g[nodeVal] = make(map[string]int)
			neighbours := c.GetNeighboursPos(false)
			for _, n := range neighbours {
				if !n.IsValid(len(grid[0]), len(grid)) {
					continue
				}
				if !grid[n.Y][n.X] {
					continue
				}
				neighVal := fmt.Sprintf("%v,%v", n.X, n.Y)
				g[nodeVal][neighVal] = 1
			}
		}
	}
	return g
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
