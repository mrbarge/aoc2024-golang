package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/albertorestifo/dijkstra"
	"github.com/mrbarge/aoc2024-golang/helper"
)

type NumpadBot struct {
	pos         helper.Coord
	numpadGraph dijkstra.Graph
	numKeypad   map[rune]helper.Coord
	numpad      [][]bool
}

type DirpadBot struct {
	pos         helper.Coord
	dirpadGraph dijkstra.Graph
	dirpad      [][]bool
	dirKeypad   map[helper.Direction]helper.Coord
}

func (n *NumpadBot) Move(r rune) []helper.Direction {
	numpadPath, _, _ := n.numpadGraph.Path(n.pos.String(), n.numKeypad[r].String())
	numpadDirections := pathToDirections(numpadPath)
	// Got to press A at the end!
	numpadDirections = append(numpadDirections, helper.NORTHEAST)
	n.pos = n.numKeypad[r]
	return numpadDirections
}

func (n *DirpadBot) Move(r helper.Direction) []helper.Direction {
	dirpadPath, _, _ := n.dirpadGraph.Path(n.pos.String(), n.dirKeypad[r].String())
	dirpadDirections := pathToDirections(dirpadPath)
	// Got to press A at the end!
	dirpadDirections = append(dirpadDirections, helper.NORTHEAST)
	n.pos = n.dirKeypad[r]
	return dirpadDirections
}

type WorldState struct {
	numpadBot  *NumpadBot
	dirpadBot  *DirpadBot
	dirpadBot2 *DirpadBot
}

func readData(lines []string) WorldState {
	w := WorldState{}
	numKeypad := map[rune]helper.Coord{
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
	numpad := [][]bool{
		{true, true, true},
		{true, true, true},
		{true, true, true},
		{false, true, true},
	}
	dirpad := [][]bool{
		{false, true, true},
		{true, true, true},
	}
	dirKeypad := map[helper.Direction]helper.Coord{
		helper.NORTH:     helper.Coord{X: 1},
		helper.NORTHEAST: helper.Coord{X: 2},
		helper.EAST:      helper.Coord{X: 2, Y: 1},
		helper.SOUTH:     helper.Coord{X: 1, Y: 1},
		helper.WEST:      helper.Coord{Y: 1},
	}

	w.numpadBot = &NumpadBot{
		pos:         helper.Coord{X: 2, Y: 3},
		numpadGraph: padGraph(numpad),
		numKeypad:   numKeypad,
		numpad:      numpad,
	}
	w.dirpadBot = &DirpadBot{
		pos:         helper.Coord{X: 2, Y: 0},
		dirpadGraph: padGraph(dirpad),
		dirpad:      dirpad,
		dirKeypad:   dirKeypad,
	}
	w.dirpadBot2 = &DirpadBot{
		pos:         helper.Coord{X: 2, Y: 0},
		dirpadGraph: padGraph(dirpad),
		dirpad:      dirpad,
		dirKeypad:   dirKeypad,
	}

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

func move(dest rune, w *WorldState) []helper.Direction {

	allDirs := make([]helper.Direction, 0)

	// Find the directions that the dirpad bot needs to push
	numpadDirs := w.numpadBot.Move(dest)

	for _, numpadDir := range numpadDirs {

		// for each direction numpad bot must move, dirpad bot must:
		// - press that direction button
		// - press 'A'

		// Find the directions that bot 2 needs to push to move bot1
		b1DirPadDirs := w.dirpadBot.Move(numpadDir)
		b1Dirs := append([]helper.Direction{}, b1DirPadDirs...)

		for _, b1Dir := range b1Dirs {
			b2DirPadDirs := w.dirpadBot2.Move(b1Dir)
			allDirs = append(allDirs, b2DirPadDirs...)
		}
	}

	return allDirs
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

func (w *WorldState) Reset() {
	w.dirpadBot2.pos = helper.Coord{X: 2}
	w.dirpadBot.pos = helper.Coord{X: 2}
	w.numpadBot.pos = helper.Coord{X: 2, Y: 3}
}

func simulate(code string, w *WorldState) int {
	allMoves := make([]helper.Direction, 0)
	for _, n := range code {
		moves := move(n, w)
		allMoves = append(allMoves, moves...)
	}
	return len(allMoves)
}

func simulateForSmallest(code string, w *WorldState) int {
	m := math.MaxInt
	for i := 0; i < 2000; i++ {
		l := simulate(code, w)
		if l < m {
			m = l
		}
		w.Reset()
	}
	return m
}

func getNumeric(code string) int {
	newCode := strings.Replace(code, "A", "", -1)
	codeNum, _ := strconv.Atoi(newCode)
	return codeNum
}

func partone(lines []string) (r int, err error) {
	worldstate := readData(lines)

	for _, line := range lines {
		pathlen := simulateForSmallest(line, &worldstate)
		r += pathlen * getNumeric(line)
	}
	return r, nil
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
