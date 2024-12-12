package main

import (
	"fmt"
	"os"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Region struct {
	Name   string
	Coords map[helper.Coord]struct{}
}

func (r Region) Area() int {
	return len(r.Coords)
}

func (r Region) Perimeter() int {
	edges := 0
	for coord, _ := range r.Coords {
		neighbours := coord.GetNeighbours(false)
		for _, neighbour := range neighbours {
			if _, ok := r.Coords[neighbour]; !ok {
				// This is an edge because it's not a neighbour
				edges++
			}
		}
	}
	return edges
}

func (r Region) NumSides() int {
	sides := 0
	for coord, _ := range r.Coords {
		neighbours := coord.GetNeighboursAsMap(true)
		shapeNeighbours := make(map[helper.Direction]bool)
		for dir, neighbour := range neighbours {
			// flag if a neighbour comprises the shape
			if _, ok := r.Coords[neighbour]; ok {
				shapeNeighbours[dir] = true
			} else {
				shapeNeighbours[dir] = false
			}
		}
		corners := make([]bool, 0)
		corners = append(corners, !shapeNeighbours[helper.NORTH] && !shapeNeighbours[helper.WEST])
		corners = append(corners, !shapeNeighbours[helper.NORTH] && !shapeNeighbours[helper.EAST])
		corners = append(corners, !shapeNeighbours[helper.SOUTH] && !shapeNeighbours[helper.WEST])
		corners = append(corners, !shapeNeighbours[helper.SOUTH] && !shapeNeighbours[helper.EAST])
		corners = append(corners, shapeNeighbours[helper.SOUTH] && shapeNeighbours[helper.EAST] && !shapeNeighbours[helper.SOUTHEAST])
		corners = append(corners, shapeNeighbours[helper.SOUTH] && shapeNeighbours[helper.WEST] && !shapeNeighbours[helper.SOUTHWEST])
		corners = append(corners, shapeNeighbours[helper.NORTH] && shapeNeighbours[helper.WEST] && !shapeNeighbours[helper.NORTHWEST])
		corners = append(corners, shapeNeighbours[helper.NORTH] && shapeNeighbours[helper.EAST] && !shapeNeighbours[helper.NORTHEAST])

		for _, corner := range corners {
			if corner {
				sides++
			}
		}
	}
	return sides
}

func readData(lines []string) [][]string {
	grid := make([][]string, 0)
	for _, line := range lines {
		row := make([]string, 0)
		for _, v := range line {
			row = append(row, string(v))
		}
		grid = append(grid, row)
	}
	return grid
}

func findRegions(grid [][]string) []*Region {
	regionMap := make(map[helper.Coord]*Region)
	regions := make([]*Region, 0)
	for y, line := range grid {
		for x, v := range line {
			coord := helper.Coord{X: x, Y: y}
			// Did we already find a region for this?
			if _, ok := regionMap[coord]; ok {
				// yeah, skip it
				continue
			}
			r := Region{
				Name:   v,
				Coords: map[helper.Coord]struct{}{coord: {}},
			}
			regions = append(regions, &r)
			regionMap[coord] = &r
			addNeighboursToRegion(r.Name, coord, regionMap, grid)
		}
	}
	return regions
}

func addNeighboursToRegion(name string, c helper.Coord, regionMap map[helper.Coord]*Region, grid [][]string) {
	neighbours := c.GetNeighbours(false)
	for _, neighbour := range neighbours {
		if !neighbour.IsValid(len(grid[0]), len(grid)) {
			// off the grid, ignore
			continue
		} else if grid[neighbour.Y][neighbour.X] != name {
			// not part of the region, ignore
			continue
		} else if _, ok := regionMap[neighbour]; ok {
			// part of some other region, ignore
			continue
		}
		region := regionMap[c]
		region.Coords[neighbour] = struct{}{}
		regionMap[neighbour] = region
		addNeighboursToRegion(name, neighbour, regionMap, grid)
	}
}

func partone(lines []string) (r int, err error) {
	grid := readData(lines)
	regions := findRegions(grid)
	for _, region := range regions {
		r += region.Perimeter() * region.Area()
	}
	return r, nil
}

func parttwo(lines []string) (r int, err error) {
	grid := readData(lines)
	regions := findRegions(grid)
	for _, region := range regions {
		r += region.NumSides() * region.Area()
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
