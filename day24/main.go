package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type GateRule struct {
	lhs    string
	op     string
	rhs    string
	dest   string
	hasRun bool
}

func readData(lines []string) (map[string]bool, []*GateRule) {
	initialValues := true

	gates := make(map[string]bool)
	gaterules := make([]*GateRule, 0)
	for _, line := range lines {
		if line == "" {
			initialValues = false
			continue
		}

		if initialValues {
			gate := strings.Split(line, ": ")[0]
			gateval, _ := strconv.Atoi(strings.Split(line, ": ")[1])
			gates[gate] = gateval == 1
		} else {
			elems := strings.Split(line, " ")
			gr := GateRule{
				lhs:  elems[0],
				op:   elems[1],
				rhs:  elems[2],
				dest: elems[4],
			}
			gaterules = append(gaterules, &gr)
		}
	}
	return gates, gaterules
}

func op(lhs bool, op string, rhs bool) bool {
	switch op {
	case "XOR":
		return lhs != rhs
	case "AND":
		return lhs && rhs
	case "OR":
		return lhs || rhs
	}
	return false
}

func run(gates map[string]bool, gaterules []*GateRule) {

	for {
		ranRule := false
		for _, rule := range gaterules {
			if rule.hasRun {
				continue
			}

			canRunRule := true
			for _, gate := range []string{rule.lhs, rule.rhs} {
				if _, ok := gates[gate]; !ok {
					canRunRule = false
				}
			}
			if !canRunRule {
				continue
			}

			gates[rule.dest] = op(gates[rule.lhs], rule.op, gates[rule.rhs])
			ranRule = true
			rule.hasRun = true
		}
		if !ranRule {
			break
		}
	}
}

func makeZArray(gates map[string]bool) int64 {
	zarray := ""
	for i := 0; i < len(gates); i++ {
		zgate := fmt.Sprintf("z%02d", i)
		if _, ok := gates[zgate]; !ok {
			break
		}
		if gates[zgate] {
			zarray = "1" + zarray
		} else {
			zarray = "0" + zarray
		}
	}
	decimal, _ := strconv.ParseInt(zarray, 2, 64)
	return decimal
}

func partone(lines []string) (r int64, err error) {
	gates, gaterules := readData(lines)
	run(gates, gaterules)
	return makeZArray(gates), nil
}

func parttwo(lines []string) (r int64, err error) {
	gates, gaterules := readData(lines)

	dests := make([]string, 0)
	for _, gaterule := range gaterules {
		if !slices.Contains(dests, gaterule.dest) {
			dests = append(dests, gaterule.dest)
		}
	}
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

	ans, err = parttwo(lines)
	fmt.Printf("Part two: %v\n", ans)

}
