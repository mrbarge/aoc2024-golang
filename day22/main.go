package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type BuyerSequence struct {
	secretNumbers []int64
	changes       []int64
}

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

func calculateBuyerSequence(in int64) BuyerSequence {
	b := BuyerSequence{
		secretNumbers: make([]int64, 0),
		changes:       make([]int64, 0),
	}

	currentSecretNumber := in
	for i := 0; i < 2000; i++ {
		b.secretNumbers = append(b.secretNumbers, currentSecretNumber)
		nextSecretNumber := step(currentSecretNumber)
		change := (nextSecretNumber % 10) - (currentSecretNumber % 10)
		b.changes = append(b.changes, change)
		currentSecretNumber = nextSecretNumber
	}
	b.secretNumbers = append(b.secretNumbers, currentSecretNumber)
	return b
}

func partone(lines []string) (r int64, err error) {
	for _, line := range lines {
		v, _ := strconv.Atoi(line)
		r += rotate(int64(v), 2000)
	}
	return r, nil
}

func makeSeenKey(n []int64) string {
	s := make([]string, 0)
	for _, v := range n {
		s = append(s, strconv.Itoa(int(v)))
	}
	return strings.Join(s, ",")
}

func parttwo(lines []string) (r int64, err error) {
	buyerSequences := make([]BuyerSequence, 0)
	for _, line := range lines {
		v, _ := strconv.Atoi(line)
		buyerSequence := calculateBuyerSequence(int64(v))
		buyerSequences = append(buyerSequences, buyerSequence)
	}

	bananaWinnings := make(map[string]int64)
	for _, buyerSequence := range buyerSequences {
		seenRanges := make(map[string]int64)
		for i := 0; i < len(buyerSequence.changes)-3; i++ {
			subrange := buyerSequence.changes[i : i+4]
			subrangeKey := makeSeenKey(subrange)
			if _, ok := seenRanges[subrangeKey]; !ok {
				seenRanges[subrangeKey] = (buyerSequence.secretNumbers[i+4]) % 10
				if _, ok := bananaWinnings[subrangeKey]; !ok {
					bananaWinnings[subrangeKey] = 0
				}
				bananaWinnings[subrangeKey] += seenRanges[subrangeKey]
			}
		}
	}
	winnings := make([]int64, 0)
	for _, v := range bananaWinnings {
		winnings = append(winnings, v)
	}
	return slices.Max(winnings), nil

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
