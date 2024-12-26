package main

import (
	"fmt"
	"os"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Schematic [][]bool
type Key Schematic
type Lock Schematic

func readData(lines []string) ([]Key, []Lock) {

	keys := make([]Key, 0)
	locks := make([]Lock, 0)

	tempgrid := make([][]bool, 0)
	for _, line := range lines {
		if line == "" {
			if isLock(tempgrid) {
				locks = append(locks, tempgrid)
			} else {
				keys = append(keys, tempgrid)
			}
			tempgrid = make([][]bool, 0)
			continue
		}
		row := make([]bool, 0)
		for _, v := range line {
			row = append(row, v == '#')
		}
		tempgrid = append(tempgrid, row)
	}
	if isLock(tempgrid) {
		locks = append(locks, tempgrid)
	} else {
		keys = append(keys, tempgrid)
	}
	return keys, locks
}

func (k Key) heights() []int {
	r := make([]int, len(k[0]))
	for y := 0; y < len(k)-1; y++ {
		for x := 0; x < len(k[y]); x++ {
			if k[y][x] {
				r[x]++
			}
		}
	}
	return r
}

func (l Lock) heights() []int {
	r := make([]int, len(l[0]))
	for y := 1; y < len(l); y++ {
		for x := 0; x < len(l[y]); x++ {
			if l[y][x] {
				r[x]++
			}
		}
	}
	return r
}

func isLock(grid [][]bool) bool {
	for i := 0; i < len(grid[0]); i++ {
		if !grid[0][i] {
			return false
		}
	}
	return true
}

func overlap(k Key, l Lock) bool {
	for i := 0; i < len(k.heights()); i++ {
		if 6-k.heights()[i]-l.heights()[i] <= 0 {
			return true
		}
	}
	return false
}

func partone(lines []string) (r int, err error) {
	keys, locks := readData(lines)
	keys = keys
	locks = locks

	for _, k := range keys {
		for _, l := range locks {
			if !overlap(k, l) {
				r++
			}
		}
	}

	return r, nil
}

func parttwo(lines []string) (r int, err error) {
	return 0, nil
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

}
