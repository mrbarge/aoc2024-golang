package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mrbarge/aoc2024-golang/helper"
)

func step(in int64) int64 {
	s1 := prune(mix(in*64, in))
	s2 := prune(mix(s1/32, s1))
	s3 := prune(mix(s2*2048, s2))
	return s3
}

func rotate(in int64, steps int) int64 {
	for i := 0; i < steps; i++ {
		in = step(in)
	}
	return in
}

func mix(in int64, secret int64) int64 {
	return in ^ secret
}

func prune(in int64) int64 {
	return in % 16777216
}

func partone(lines []string) (r int64, err error) {
	for _, line := range lines {
		v, _ := strconv.Atoi(line)
		r += rotate(int64(v), 2000)
	}
	return r, nil
}

func parttwo(lines []string) (r int64, err error) {
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
