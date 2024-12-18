package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/albertorestifo/dijkstra"
	"github.com/mrbarge/aoc2024-golang/helper"
)

func readData(lines []string) []helper.Coord {
	r := make([]helper.Coord, 0)
	for _, line := range lines {
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		r = append(r, helper.Coord{X: x, Y: y})
	}
	return r
}

func makeGrid(sizex int, sizey int) [][]bool {
	r := make([][]bool, sizey)
	for y := 0; y < sizey; y++ {
		r[y] = make([]bool, sizex)
	}
	return r
}

func readData2(lines []string, sizex int, sizey int) [][]bool {
	r := make([][]bool, sizey)
	for y := 0; y < sizey; y++ {
		r[y] = make([]bool, sizex)
	}
	for _, line := range lines {
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		r[y][x] = true
	}
	return r
}

func fall(grid [][]bool) [][]bool {
	r := make([][]bool, len(grid))
	for y, row := range grid {
		r[y] = make([]bool, len(row))
	}

	for y, row := range grid {
		for x, v := range row {
			if v {
				r[(y+1)%len(grid)][x] = true
			}
		}
	}
	return r
}

func printGrid(grid [][]bool) {
	for _, row := range grid {
		for _, v := range row {
			if v {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func shortest(grid [][]bool) ([]string, int) {
	g := dijkstra.Graph{}

	for y, row := range grid {
		for x, v := range row {
			if v {
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
				if grid[n.Y][n.X] {
					continue
				}
				neighVal := fmt.Sprintf("%v,%v", n.X, n.Y)
				g[nodeVal][neighVal] = 1
			}
		}
	}

	end := fmt.Sprintf("%v,%v", len(grid[0])-1, len(grid)-1)
	path, score, _ := g.Path("0,0", end)
	return path, score
}

func partone(lines []string) (r int, err error) {
	//sizex, sizey := 7, 7
	//ticks := 12
	sizex, sizey := 71, 71
	ticks := 1024

	grid := makeGrid(sizex, sizey)
	bs := readData(lines)
	for i := 0; i < ticks; i++ {
		c := bs[i]
		grid[c.Y][c.X] = true
	}

	_, s := shortest(grid)
	return s, nil
}

func parttwo(lines []string) (r int, err error) {
	//sizex, sizey := 7, 7
	sizex, sizey := 71, 71

	grid := makeGrid(sizex, sizey)
	bs := readData(lines)
	var nc helper.Coord
	for i := 0; i < len(lines); i++ {
		c := bs[i]
		grid[c.Y][c.X] = true
		nc = c
		_, s := shortest(grid)
		if s == 0 {
			break
		}
	}
	fmt.Printf("%v\n", nc)
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
