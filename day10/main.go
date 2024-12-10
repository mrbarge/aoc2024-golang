package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type TrailMap struct {
	Width  int
	Height int
	Grid   [][]int
}

type SuccessfulTrail struct {
	Walked []helper.Coord
	Ending helper.Coord
}

func readData(lines []string) TrailMap {
	t := TrailMap{
		Width:  len(lines[0]),
		Height: len(lines),
		Grid:   make([][]int, 0),
	}
	for _, line := range lines {
		row := make([]int, 0)
		for _, v := range line {
			i, _ := strconv.Atoi(string(v))
			row = append(row, i)
		}
		t.Grid = append(t.Grid, row)
	}
	return t
}

func (t TrailMap) trailheads() (r []helper.Coord) {
	r = make([]helper.Coord, 0)
	for y, line := range t.Grid {
		for x, v := range line {
			if v == 0 {
				r = append(r, helper.Coord{X: x, Y: y})
			}
		}
	}
	return r
}

func walkTrail(pos helper.Coord, walked []helper.Coord, level int, t TrailMap) []SuccessfulTrail {
	successfulTrails := make([]SuccessfulTrail, 0)

	if !pos.IsValid(t.Width, t.Height) {
		// We walked off the map, shouldn't happen but whatever
		return []SuccessfulTrail{}
	}
	if t.Grid[pos.Y][pos.X] == 9 {
		// We found a good trail
		return []SuccessfulTrail{{Walked: append(walked, pos), Ending: pos}}
	}

	neighbours := pos.GetNeighbours(false)
	for _, neighbour := range neighbours {
		if !neighbour.IsValid(t.Width, t.Height) {
			continue
		}
		if t.Grid[neighbour.Y][neighbour.X] != level+1 {
			continue
		}
		trails := walkTrail(neighbour, append(walked, pos), level+1, t)
		successfulTrails = append(successfulTrails, trails...)
	}
	return successfulTrails
}

func uniquePeaks(trails []SuccessfulTrail) int {
	m := make(map[helper.Coord]bool)
	for _, trail := range trails {
		m[trail.Walked[len(trail.Walked)-1]] = true
	}
	return len(m)
}

func partone(lines []string) (r int, err error) {
	trailmap := readData(lines)
	scores := make(map[helper.Coord]int)

	for _, trailhead := range trailmap.trailheads() {
		trails := walkTrail(trailhead, []helper.Coord{}, 0, trailmap)
		scores[trailhead] = uniquePeaks(trails)
	}

	for _, score := range scores {
		r += score
	}
	return r, nil
}

func parttwo(lines []string) (r int, err error) {
	trailmap := readData(lines)
	scores := make(map[helper.Coord]int)

	for _, trailhead := range trailmap.trailheads() {
		trails := walkTrail(trailhead, []helper.Coord{}, 0, trailmap)
		scores[trailhead] = len(trails)
	}

	for _, score := range scores {
		r += score
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
