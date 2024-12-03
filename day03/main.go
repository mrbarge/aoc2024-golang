package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type ParseState int

const (
	MUL_SEARCH ParseState = iota
	MUL_START
	MUL_FIRST
	MUL_SECOND
	MUL_END
)

func scan(s string, part_two bool) (r int) {
	idx := 0
	state := MUL_SEARCH
	numbers := make(map[ParseState]string)
	enabled := true
outer:
	for idx < len(s) {
		switch state {
		case MUL_SEARCH:
			// We're looking for the next 'mul'
			numbers[MUL_FIRST] = ""
			numbers[MUL_SECOND] = ""

			mul_start := strings.Index(s[idx:], "mul(")
			mul_do := strings.Index(s[idx:], "do()")
			mul_dont := strings.Index(s[idx:], "don't()")

			if mul_start < 0 {
				break outer
			}

			if mul_do >= 0 && (mul_dont < 0 || mul_dont > mul_do) && mul_do < mul_start {
				enabled = true
				idx += mul_do + 4
			} else if mul_dont >= 0 && (mul_do < 0 || mul_do > mul_dont) && mul_dont < mul_start {
				enabled = false
				idx += mul_dont + 7
			} else {
				state = MUL_FIRST
				idx += mul_start + 4
			}

		case MUL_FIRST, MUL_SECOND:
			if s[idx] >= '0' && s[idx] <= '9' {
				numbers[state] += string(s[idx])
			} else if state == MUL_FIRST && s[idx] == ',' {
				state = MUL_SECOND
			} else if state == MUL_SECOND && s[idx] == ')' {
				state = MUL_END
			} else {
				state = MUL_SEARCH
			}
			idx++
		case MUL_END:
			d1, err1 := strconv.Atoi(numbers[MUL_FIRST])
			d2, err2 := strconv.Atoi(numbers[MUL_SECOND])
			if err1 == nil && err2 == nil {
				if !part_two || enabled {
					r += d1 * d2
				}
			}
			state = MUL_SEARCH
		default:
			break outer
		}
	}
	return r
}

func partone(lines []string) (r int, err error) {
	answer := scan(lines[0], false)
	return answer, nil
}

func parttwo(lines []string) (r int, err error) {
	answer := scan(lines[0], true)
	return answer, nil
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
