package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Robot struct {
	Pos helper.Coord
	Vel helper.Coord
}

func readData(lines []string) []Robot {
	r := make([]Robot, 0)
	for _, line := range lines {
		rawpos := strings.Split(line, " ")[0]
		rawvel := strings.Split(line, " ")[1]
		pos := helper.ParseCoord(strings.Split(rawpos, "=")[1])
		vel := helper.ParseCoord(strings.Split(rawvel, "=")[1])
		r = append(r, Robot{Pos: pos, Vel: vel})
	}
	return r
}

func tick(robots []Robot, sizex int, sizey int) []Robot {
	r := make([]Robot, 0)
	for _, robot := range robots {
		newrobot := Robot{
			Pos: robot.Pos.MoveGridWithVelocity(robot.Vel, sizex, sizey),
			Vel: robot.Vel,
		}
		r = append(r, newrobot)
	}
	return r
}

func printGrid(robots []Robot, sizex int, sizey int) {
	grid := make([][]bool, 0)
	for i := 0; i < sizey; i++ {
		grid = append(grid, make([]bool, sizex))
	}
	for _, robot := range robots {
		grid[robot.Pos.Y][robot.Pos.X] = true
	}
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}

func findTree(robots []Robot, sizex int, sizey int) bool {
	grid := make([][]bool, 0)
	for i := 0; i < sizey; i++ {
		grid = append(grid, make([]bool, sizex))
	}
	for _, robot := range robots {
		grid[robot.Pos.Y][robot.Pos.X] = true
	}
	for _, row := range grid {
		// look for a row with lots of robots in it
		// probably means there's a tree maybe?
		rowCount := 0
		for _, cell := range row {
			if cell {
				rowCount++
			}
		}
		if rowCount >= (sizex / 3) {
			return true
		}
	}
	return false
}

func partone(lines []string) (r int, err error) {
	sizex := 101
	sizey := 103

	//sizex := 11
	//sizey := 7

	ticks := 100
	robots := readData(lines)
	for _ = range ticks {
		robots = tick(robots, sizex, sizey)
	}
	quadrants := make([]int, 4)
	for _, robot := range robots {
		if robot.Pos.X == (sizex/2) ||
			robot.Pos.Y == (sizey/2) {
			// Ignore
			continue
		}
		if robot.Pos.X < (sizex / 2) {
			// either quadrant 0 or 2
			if robot.Pos.Y < (sizey / 2) {
				quadrants[0]++
			} else {
				quadrants[2]++
			}
		} else {
			// either quadrant 1 or 3
			if robot.Pos.Y < (sizey / 2) {
				quadrants[1]++
			} else {
				quadrants[3]++
			}
		}
	}
	safety := 1
	for _, v := range quadrants {
		safety *= v
	}
	return safety, nil
}

func parttwo(lines []string) (r int, err error) {
	sizex := 101
	sizey := 103

	//sizex := 11
	//sizey := 7

	robots := readData(lines)
	for t := 0; ; t++ {
		robots = tick(robots, sizex, sizey)
		if findTree(robots, sizex, sizey) {
			printGrid(robots, sizex, sizey)
			return t, nil
		}
	}
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
