package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/albertorestifo/dijkstra"
	"github.com/mrbarge/aoc2024-golang/helper"
)

type SeenSequence struct {
	sequence  []helper.Direction
	iteration int
}

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

type WorldState struct {
	numpadBot  *NumpadBot
	dirpadBot  *DirpadBot
	dirpadBot2 *DirpadBot

	numKeypad map[rune]helper.Coord
	numpad    [][]bool
	dirpad    [][]bool
	dirKeypad map[helper.Direction]helper.Coord

	numpadPaths map[rune]map[rune][]helper.Direction
	dirpadPaths map[helper.Direction]map[helper.Direction][]helper.Direction
}

func (w *WorldState) PrecalculateNumpad() map[rune]map[rune][]helper.Direction {
	values := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A'}

	r := make(map[rune]map[rune][]helper.Direction)
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values); j++ {
			from := values[i]
			to := values[j]
			shortest := w.numpadBot.MoveToFrom(from, to)
			if _, ok := r[from]; !ok {
				r[from] = make(map[rune][]helper.Direction)
			}
			r[from][to] = shortest
		}
	}
	return r
}

func (w *WorldState) PrecalculateDirpad() map[helper.Direction]map[helper.Direction][]helper.Direction {
	values := []helper.Direction{helper.WEST, helper.NORTHEAST, helper.NORTH, helper.SOUTH, helper.EAST}

	r := make(map[helper.Direction]map[helper.Direction][]helper.Direction)
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values); j++ {
			from := values[i]
			to := values[j]
			shortest := w.dirpadBot.MoveToFrom(from, to)
			if _, ok := r[from]; !ok {
				r[from] = make(map[helper.Direction][]helper.Direction)
			}
			r[from][to] = shortest
		}
	}
	return r
}

func (n *NumpadBot) MoveToFrom(from rune, to rune) []helper.Direction {
	fromPos := n.numKeypad[from]
	toPos := n.numKeypad[to]
	numpadPath, _, _ := n.numpadGraph.Path(fromPos.String(), toPos.String())
	numpadDirections := pathToDirections(numpadPath)
	// Got to press A at the end!
	numpadDirections = append(numpadDirections, helper.NORTHEAST)
	return numpadDirections
}

func (n *DirpadBot) MoveToFrom(from helper.Direction, to helper.Direction) []helper.Direction {
	fromPos := n.dirKeypad[from]
	toPos := n.dirKeypad[to]
	numpadPath, _, _ := n.dirpadGraph.Path(fromPos.String(), toPos.String())
	numpadDirections := pathToDirections(numpadPath)
	// Got to press A at the end!
	numpadDirections = append(numpadDirections, helper.NORTHEAST)
	return numpadDirections
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

func optimizePath(d []helper.Direction) []helper.Direction {
	order := map[helper.Direction]int{
		helper.WEST:      0,
		helper.SOUTH:     1,
		helper.NORTH:     2,
		helper.EAST:      3,
		helper.NORTHEAST: 4,
	}

	// Sort the slice using the custom order
	sort.Slice(d, func(i, j int) bool {
		return order[d[i]] < order[d[j]]
	})

	return d
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

	w.numpadBot = &NumpadBot{
		pos:         helper.Coord{X: 2, Y: 3},
		numpadGraph: padGraph(w.numpad),
		numKeypad:   w.numKeypad,
		numpad:      w.numpad,
	}
	w.dirpadBot = &DirpadBot{
		pos:         helper.Coord{X: 2, Y: 0},
		dirpadGraph: padGraph(w.dirpad),
		dirpad:      w.dirpad,
		dirKeypad:   w.dirKeypad,
	}
	w.dirpadBot2 = &DirpadBot{
		pos:         helper.Coord{X: 2, Y: 0},
		dirpadGraph: padGraph(w.dirpad),
		dirpad:      w.dirpad,
		dirKeypad:   w.dirKeypad,
	}
	w.numpadPaths = w.PrecalculateNumpad()
	w.dirpadPaths = w.PrecalculateDirpad()

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

func numPadLength(code string, robots int, w *WorldState) int {
	numpadBot := 'A'
	length := 0
	seen := make(map[string]int)
	for _, nextButton := range code {
		dirpath := w.numpadPaths[numpadBot][nextButton]
		res := dirPathLength(dirpath, robots, seen, w)
		length += res
		numpadBot = nextButton
	}
	return length
}

func makeKey(dirs []helper.Direction, robot int) string {
	r := fmt.Sprintf("%v:", robot)
	for _, v := range dirs {
		r += fmt.Sprintf("%v", v)
	}
	return r
}

func dirPathLength(dirs []helper.Direction, robot int, seen map[string]int, w *WorldState) int {
	if robot == 0 {
		return len(dirs)
	}
	if _, ok := seen[makeKey(dirs, robot)]; ok {
		return seen[makeKey(dirs, robot)]
	}

	length := 0
	// All bots start at A
	var dirpadBot helper.Direction = helper.NORTHEAST
	for _, nextDir := range dirs {
		dirpath := w.dirpadPaths[dirpadBot][nextDir]
		res := dirPathLength(dirpath, robot-1, seen, w)
		length += res
		dirpadBot = nextDir
	}
	seen[makeKey(dirs, robot)] = length
	return length
}

func parttwo(lines []string) (r int, err error) {
	for _, line := range lines {
		pathlen := math.MaxInt
		for i := 0; i < 5000; i++ {
			worldstate := readData(lines)
			m := numPadLength(line, 25, &worldstate)
			if m < pathlen {
				pathlen = m
			}
		}
		r += pathlen * getNumeric(line)
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
