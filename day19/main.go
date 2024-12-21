package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrbarge/aoc2024-golang/helper"
)

func readData(lines []string) (patterns []string, designs []string) {
	patterns = strings.Split(lines[0], ", ")
	designs = make([]string, 0)

	for _, line := range lines[1:] {
		designs = append(designs, line)
	}

	return patterns, designs
}

func isValid(design string, patterns []string, cache map[string]int) (awoo int) {
	if len(design) == 0 {
		cache[design] = 1
		return 1
	}

	if _, ok := cache[design]; ok {
		return cache[design]
	}

	awoo = 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			validCombinations := isValid(design[len(pattern):], patterns, cache)
			cache[design[len(pattern):]] = validCombinations
			awoo += validCombinations
		}
	}
	return awoo
}

func partone(lines []string) (yarr int, err error) {
	patterns, designs := readData(lines)

	for _, design := range designs {
		cache := make(map[string]int)
		if isValid(design, patterns, cache) > 0 {
			yarr++
		}
	}
	return yarr, nil
}

func parttwo(lines []string) (yarr int, err error) {
	patterns, designs := readData(lines)

	for _, design := range designs {
		cache := make(map[string]int)
		yarr += isValid(design, patterns, cache)
	}
	return yarr, nil
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
