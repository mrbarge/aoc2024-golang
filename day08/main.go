package main

import (
	"fmt"
	"maps"
	"os"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Node struct {
	freq string
	c    helper.Coord
}

func readData(lines []string) (r map[string][]Node) {
	r = make(map[string][]Node)
	for y, line := range lines {
		for x, v := range line {
			if v == '.' {
				continue
			}
			freq := string(v)
			if _, ok := r[freq]; !ok {
				r[freq] = make([]Node, 0)
			}
			r[freq] = append(r[freq], Node{
				freq: freq,
				c:    helper.Coord{X: x, Y: y},
			})
		}
	}
	return r
}

func makePairs(nodes []Node) [][2]Node {
	r := make([][2]Node, 0)
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			r = append(r, [2]Node{nodes[i], nodes[j]})
		}
	}
	return r
}

func makeAntinodes(c1 helper.Coord, c2 helper.Coord) [2]helper.Coord {
	return [2]helper.Coord{
		{
			X: c1.X + (c1.X - c2.X),
			Y: c1.Y + (c1.Y - c2.Y),
		},
		{
			X: c2.X + (c2.X - c1.X),
			Y: c2.Y + (c2.Y - c1.Y),
		},
	}
}

func makeSuperAntinodes(c1 helper.Coord, c2 helper.Coord, maxX int, maxY int) []helper.Coord {
	antinodeMap := make(map[helper.Coord]bool)
	for _, pair := range [][2]helper.Coord{{c1, c2}, {c2, c1}} {
		p1 := pair[0]
		p2 := pair[1]
		for {
			testynode := helper.Coord{
				X: p1.X + (p1.X - p2.X),
				Y: p1.Y + (p1.Y - p2.Y),
			}
			if testynode.IsValid(maxX, maxY) {
				antinodeMap[testynode] = true
			} else {
				break
			}
			p2 = p1
			p1 = testynode
		}
	}
	antinodes := make([]helper.Coord, 0)
	for a := range maps.Keys(antinodeMap) {
		antinodes = append(antinodes, a)
	}
	return antinodes
}

func partone(lines []string) (r int, err error) {
	antimap := make(map[helper.Coord]bool)
	freqs := readData(lines)
	for _, nodes := range freqs {
		pairs := makePairs(nodes)
		for _, pair := range pairs {
			antinodes := makeAntinodes(pair[0].c, pair[1].c)
			for _, antinode := range antinodes {
				if antinode.IsValid(len(lines[0]), len(lines)) {
					antimap[antinode] = true
				}
			}
		}
	}
	return len(antimap), nil

}

func parttwo(lines []string) (r int, err error) {
	antimap := make(map[helper.Coord]bool)
	freqs := readData(lines)
	for _, nodes := range freqs {
		for _, node := range nodes {
			antimap[node.c] = true
		}
		pairs := makePairs(nodes)
		for _, pair := range pairs {
			antinodes := makeSuperAntinodes(pair[0].c, pair[1].c, len(lines[0]), len(lines))
			for _, antinode := range antinodes {
				antimap[antinode] = true
			}
		}
	}
	return len(antimap), nil
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
