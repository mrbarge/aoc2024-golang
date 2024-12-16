package main

import (
	"fmt"
	"os"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Box struct {
	left  helper.Coord
	right helper.Coord
}

type GameState struct {
	robot        helper.Coord
	boxmap       map[helper.Coord]bool
	boxes        []*Box
	grid         [][]bool
	instructions []helper.Direction
	tick         int
}

func readData(lines []string, partTwo bool) GameState {
	gridMode := true
	grid := make([][]bool, 0)
	boxmap := make(map[helper.Coord]bool)
	boxes := make([]*Box, 0)
	var player helper.Coord
	instructions := make([]helper.Direction, 0)
	for y, line := range lines {
		if line == "" {
			gridMode = false
			continue
		}

		row := make([]bool, 0)
		if gridMode {
			xcounter := 0
			for _, v := range line {
				c := helper.Coord{X: xcounter, Y: y}
				if v == '#' {
					row = append(row, false)
					if partTwo {
						row = append(row, false)
					}
				} else if v == 'O' {
					boxmap[c] = true
					row = append(row, true)
					if partTwo {
						c2 := helper.Coord{X: xcounter + 1, Y: y}
						box := Box{left: c, right: c2}
						boxes = append(boxes, &box)
						row = append(row, true)
					}
				} else if v == '@' {
					player = helper.Coord{X: xcounter, Y: y}
					row = append(row, true)
					if partTwo {
						row = append(row, true)
					}
				} else if v == '.' {
					row = append(row, true)
					if partTwo {
						row = append(row, true)
					}
				}
				xcounter++
				if partTwo {
					xcounter++
				}
			}
			grid = append(grid, row)
		} else {
			for _, v := range line {
				switch v {
				case '^':
					instructions = append(instructions, helper.NORTH)
				case '>':
					instructions = append(instructions, helper.EAST)
				case 'v':
					instructions = append(instructions, helper.SOUTH)
				case '<':
					instructions = append(instructions, helper.WEST)
				}
			}
		}
	}
	return GameState{
		robot:        player,
		boxmap:       boxmap,
		boxes:        boxes,
		grid:         grid,
		instructions: instructions,
		tick:         0,
	}
}

func (g *GameState) move() bool {
	if g.tick >= len(g.instructions) {
		return false
	}
	instruction := g.instructions[g.tick]
	canMove, movePos := g.canMove(instruction)
	if canMove {
		g.moveFromPos(*movePos, instruction.Opposite())
	}
	g.tick++
	return true
}

func (g *GameState) moveFromPos(startPos helper.Coord, d helper.Direction) {
	currPos := startPos
	for {
		// What's beside us?
		nextPos := currPos.Move(d)
		if nextPos == g.robot {
			// The roobot's beside us, this is the end.
			// Move the robot to where we are
			g.robot = currPos
			break
		}
		// There must be a box beside us
		if _, ok := g.boxmap[nextPos]; ok {
			// move a box
			g.boxmap[currPos] = true
			delete(g.boxmap, nextPos)
		} else {
			fmt.Printf("something odd happened..\n")
		}
		currPos = nextPos
	}
}

func (g *GameState) canMove(d helper.Direction) (bool, *helper.Coord) {
	// there needs to be at least one free space before a wall
	currPos := g.robot
	for {
		nextPos := currPos.Move(d)
		if !nextPos.IsValid(len(g.grid[0]), len(g.grid)) {
			// left the map somehow
			fmt.Printf("left the map.. yikes\n")
			return false, nil
		}
		if !g.grid[nextPos.Y][nextPos.X] {
			// we hit a wall
			return false, nil
		}
		if _, ok := g.boxmap[nextPos]; !ok {
			if g.grid[nextPos.Y][nextPos.X] {
				// found a free space
				return true, &nextPos
			}
		}
		// skip over boxes, move along
		currPos = nextPos
	}
}

func (g *GameState) MoveTwo(curPos helper.Coord, d helper.Direction) bool {
	nextPos := curPos.Move(d)

	if !nextPos.IsValid(len(g.grid[0]), len(g.grid)) { // left the map
		return false
	}
	if !g.grid[nextPos.Y][nextPos.X] { // hit a wall, can't move
		return false
	}
	hasBox := g.coordHasBox(nextPos)
	if hasBox == nil {
		// We're moving the robot to an empty space rather than pushing boxes
		g.robot = nextPos
		return true
	}
	// Robot needs to start pushin'
	result := g.MoveBox(hasBox, d)
	if result {
		g.robot = g.robot.Move(d)
		return true
	} else {
		return false
	}
}

func (b *Box) Move(d helper.Direction) {
	b.left = b.left.Move(d)
	b.right = b.right.Move(d)
}

func (g *GameState) CanMoveBox(b *Box, d helper.Direction) bool {
	nextL := b.left.Move(d)
	nextR := b.right.Move(d)
	if !g.grid[nextL.Y][nextL.X] || !g.grid[nextR.Y][nextR.X] {
		// it's a wall
		return false
	}
	return true
}

func (g *GameState) MoveBox(b *Box, d helper.Direction) bool {
	// Is there an empty state next to this box?
	if d == helper.WEST || d == helper.EAST {
		// moving sideways so this is easier
		var nextPos helper.Coord
		if d == helper.WEST {
			nextPos = b.left.Move(d)
		} else {
			nextPos = b.right.Move(d)
		}
		if !g.grid[nextPos.Y][nextPos.X] {
			// Wall in the way, can't move
			return false
		}
		hasBox := g.coordHasBox(nextPos)
		if hasBox == nil {
			// We can move
			b.Move(d)
			return true
		} else {
			// A box is in the way
			if g.MoveBox(hasBox, d) {
				// But we moved it
				b.Move(d)
				return true
			}
			// We couldn't move it
			return false
		}
	} else {
		// Moving vertically
		if !g.CanMoveBox(b, d) {
			// But we can't
			return false
		}
		connectingBoxes := g.getConnectingBoxes(b, d)
		allSuccess := true
		for _, box := range connectingBoxes {
			if !g.MoveBox(box, d) {
				allSuccess = false
			}
		}
		if allSuccess {
			// Every box could be moved, so we can move this too
			b.Move(d)
			return true
		} else {
			return false
		}
	}
}

func (g *GameState) getConnectingBoxes(b *Box, d helper.Direction) []*Box {
	nextPosL := b.left.Move(d)
	nextPosR := b.right.Move(d)
	hasBoxL := g.coordHasBox(nextPosL)
	hasBoxR := g.coordHasBox(nextPosR)
	if hasBoxL == nil && hasBoxR == nil {
		// No boxes..
		return []*Box{}
	} else if hasBoxL == hasBoxR {
		// Only one box, and it's aligned with ours
		return []*Box{hasBoxL}
	} else {
		// Two boxes
		r := make([]*Box, 0)
		if hasBoxL != nil {
			r = append(r, hasBoxL)
		}
		if hasBoxR != nil {
			r = append(r, hasBoxR)
		}
		return r
	}
}

func (g *GameState) coordHasBox(c helper.Coord) *Box {
	for _, box := range g.boxes {
		if c == box.left || c == box.right {
			return box
		}
	}
	return nil
}

func (g *GameState) Print(partTwo bool) {
	for y, row := range g.grid {
		for x, v := range row {
			c := helper.Coord{X: x, Y: y}
			hasBox := g.coordHasBox(c)
			if !v {
				fmt.Printf("#")
			} else if c == g.robot {
				fmt.Printf("@")
			} else if hasBox != nil {
				if c == hasBox.left {
					fmt.Printf("[")
				} else {
					fmt.Printf("]")
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func partone(lines []string) (r int, err error) {
	gamestate := readData(lines, false)
	for {
		if !gamestate.move() {
			break
		}
	}
	for box, _ := range gamestate.boxmap {
		r += 100*box.Y + box.X
	}
	return r, nil
}

func parttwo(lines []string) (r int, err error) {
	gamestate := readData(lines, true)
	//gamestate.Print(true)
	for i := 0; i < len(gamestate.instructions); i++ {
		d := gamestate.instructions[i]
		gamestate.MoveTwo(gamestate.robot, d)
	}
	gamestate.Print(true)
	for _, box := range gamestate.boxes {
		r += 100*box.left.Y + box.left.X
	}
	return r, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partone(lines)
	fmt.Printf("Part one: %v\n", ans)

	ans, err = parttwo(lines)
	fmt.Printf("Part two: %v\n", ans)

}
