package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Computer struct {
	id    string
	links []*Computer
}

func (c *Computer) Links(id string) bool {
	for _, v := range c.links {
		if v.id == id {
			return true
		}
	}
	return false
}

func (c *Computer) WhatLinks() []string {
	r := make([]string, 0)
	for _, v := range c.links {
		r = append(r, v.id)
	}
	sort.Strings(r)
	return r
}

func IsInChain(c *Computer, links []*Computer, computers map[string]*Computer) bool {
	for _, v := range links {
		vlinks := computers[v.id].WhatLinks()
		if slices.Contains(vlinks, c.id) {
			return true
		}
	}
	return false
}

func GetMaxInterconnected(check []*Computer, connections []*Computer, maxcon []*Computer, computers map[string]*Computer) {
	if len(connections) > len(maxcon) {
		maxcon = append([]*Computer{}, connections...)
	}
	for i, computer := range check {
		if IsInChain(computer, connections, computers) {
			GetMaxInterconnected(check[i+1:], append(connections, computer), maxcon, computers)
		}
	}
}

func GetInterconnected(computers map[string]*Computer) map[string]bool {
	r := make(map[string]bool)
	for _, c := range computers {
		for _, cn := range c.links {
			for _, cnn := range cn.links {
				if cnn.Links(c.id) {
					ids := []string{c.id, cn.id, cnn.id}
					sort.Strings(ids)
					idkey := strings.Join(ids, ",")
					r[idkey] = true
				}
			}
		}
	}
	return r
}

func readData(lines []string) map[string]*Computer {
	r := make(map[string]*Computer)
	for _, line := range lines {
		c1 := strings.Split(line, "-")[0]
		c2 := strings.Split(line, "-")[1]

		for _, v := range []string{c1, c2} {
			if _, ok := r[v]; !ok {
				r[v] = &Computer{
					id:    v,
					links: make([]*Computer, 0),
				}
			}
		}
		r[c1].links = append(r[c1].links, r[c2])
		r[c2].links = append(r[c2].links, r[c1])
	}
	return r
}

func partone(lines []string) (r int, err error) {
	computers := readData(lines)
	interconnected := GetInterconnected(computers)
	for v, _ := range interconnected {
		cids := strings.Split(v, ",")
		for _, cid := range cids {
			if strings.HasPrefix(cid, "t") {
				r++
				break
			}
		}
	}
	return r, nil

}

func parttwo(lines []string) (r int, err error) {
	computers := readData(lines)

	maxConnected := make([]*Computer, 0)
	comps := make([]*Computer, 0)
	for _, v := range computers {
		comps = append(comps, v)
	}
	GetMaxInterconnected(comps, []*Computer{}, maxConnected, computers)
	fmt.Println(maxConnected)
	return 0, nil
}

func main() {
	fh, _ := os.Open("test.txt")
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
