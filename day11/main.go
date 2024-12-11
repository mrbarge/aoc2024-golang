package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2024-golang/helper"
)

func readStonesIntoMap(line string) (r map[int]int) {
	r = make(map[int]int)
	sstones := strings.Split(line, " ")
	for _, s := range sstones {
		i, _ := strconv.Atoi(s)
		if _, ok := r[i]; !ok {
			r[i] = 1
		} else {
			r[i]++
		}
	}
	return r
}

func shiftStone(to int, total int, stones map[int]int) {
	if _, ok := stones[to]; !ok {
		stones[to] = total
	} else {
		stones[to] += total
	}
}

func tickmap(stones map[int]int) map[int]int {
	newmap := make(map[int]int)
	for stone, stoneCount := range stones {
		ss := strconv.Itoa(stone)
		if stone == 0 {
			shiftStone(1, stoneCount, newmap)
		} else if len(ss)%2 == 0 {
			// split
			n1, _ := strconv.Atoi(ss[:len(ss)/2])
			n2, _ := strconv.Atoi(ss[len(ss)/2:])
			shiftStone(n1, stoneCount, newmap)
			shiftStone(n2, stoneCount, newmap)
		} else {
			shiftStone(stone*2024, stoneCount, newmap)
		}
	}
	return newmap
}

func partone(lines []string) (r int, err error) {
	stones := readStonesIntoMap(lines[0])
	for i := 0; i < 25; i++ {
		stones = tickmap(stones)
	}
	for _, v := range stones {
		r += v
	}
	return r, nil
}

func parttwo(lines []string) (r int, err error) {
	stones := readStonesIntoMap(lines[0])
	for i := 0; i < 75; i++ {
		stones = tickmap(stones)
	}
	for _, v := range stones {
		r += v
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
