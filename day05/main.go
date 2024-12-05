package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Page struct {
	Id          int
	SmallerThan []*Page
	BiggerThan  []*Page
}

func (p *Page) Before(pageId int, seen []int) bool {
	if p.Id == pageId {
		return true
	}
	if slices.Contains(seen, p.Id) {
		return false
	}
	seen = append(seen, p.Id)
	for _, v := range p.SmallerThan {
		if v.Before(pageId, seen) {
			fmt.Printf("Found %v in %v:Before\n", v.Id, p.Id)
			return true
		}
	}
	return false
}

func (p *Page) Before2(pageId int, seen []int) bool {
	if slices.Contains(seen, p.Id) {
		return false
	}
	for _, v := range p.SmallerThan {
		// We have a rule for this page that indicates it is smaller than pageId
		if v.Id == pageId {
			return true
		}
	}
	for _, v := range p.BiggerThan {
		// We have a rule for this page that indicates it is bigger than pageId
		if v.Id == pageId {
			return false
		}
	}
	// Check all pages that this page is bigger than and make sure there is not
	// a contradicting rule for this pageId
	// A contradicting page is one that exists in v's SmallerThan list
	for _, v := range p.BiggerThan {
		if v.IsBiggerThan(pageId, []int{}) {
			return false
		}
	}
	return true
}

func (p *Page) IsBiggerThan(pageId int, seen []int) bool {
	seen = append(seen, p.Id)
	for _, v := range p.SmallerThan {
		if v.Id == pageId {
			return false
		}
		if !v.IsBiggerThan(pageId, seen) {
			return false
		}
	}
	return true
}

func (p *Page) After(pageId int, seen []int) bool {
	if p.Id == pageId {
		return true
	}
	if slices.Contains(seen, p.Id) {
		return false
	}
	seen = append(seen, p.Id)
	for _, v := range p.BiggerThan {
		if v.After(pageId, seen) {
			//fmt.Printf("Found %v in %v:After\n", pageId, p.Id)
			return true
		}
	}
	return false
}

func read_data(lines []string) (map[int]*Page, [][]int) {
	updates := make([][]int, 0)
	pages := make(map[int]*Page)

	for _, line := range lines {
		if strings.Contains(line, "|") {
			// rule
			p1, _ := strconv.Atoi(strings.Split(line, "|")[0])
			p2, _ := strconv.Atoi(strings.Split(line, "|")[1])

			for _, p := range []int{p1, p2} {
				if _, ok := pages[p]; !ok {
					pages[p] = &Page{
						Id:          p,
						SmallerThan: make([]*Page, 0),
						BiggerThan:  make([]*Page, 0),
					}
				}
			}
			pages[p1].SmallerThan = append(pages[p1].SmallerThan, pages[p2])
			pages[p2].BiggerThan = append(pages[p2].BiggerThan, pages[p1])

		} else if strings.Contains(line, ",") {
			// updates
			order := strings.Split(line, ",")
			update := make([]int, 0)
			for _, n := range order {
				snum, _ := strconv.Atoi(n)
				update = append(update, snum)
			}
			updates = append(updates, update)
		} else {
			// skip
			continue
		}
	}
	return pages, updates
}

func validate_update(pages map[int]*Page, update []int) (bool, int) {
	for i, pageId := range update[:len(update)-1] {
		nextPage := update[i+1]
		if !pages[pageId].Before2(pages[nextPage].Id, []int{}) {
			return false, i
		}
	}
	return true, 0
}

func fix_that_update(pages map[int]*Page, update []int, bad_step int) int {
	fixed := false
	test_it := make([]int, len(update))
	copy(test_it, update)

	swap_idx := bad_step
	for !fixed {
		// swap the bad index for the one beside it
		if swap_idx >= len(test_it)-1 {
			//fmt.Printf("We couldn't find a solution so this is very bad\n")
			break
		}
		tmp := test_it[swap_idx]
		test_it[swap_idx] = test_it[swap_idx+1]
		test_it[swap_idx+1] = tmp

		if ok, _ := validate_update(pages, test_it); ok {
			// It worked!!
			return test_it[len(test_it)/2]
		}
		swap_idx++
	}
	if ok, badPage := validate_update(pages, test_it); !ok {
		// Imma still fix you tho!
		return fix_that_update(pages, test_it, badPage)
	}
	return -1
}

func partone(lines []string) (r int, err error) {
	pages, updates := read_data(lines)

	for _, update := range updates {
		if ok, _ := validate_update(pages, update); ok {
			middle := update[len(update)/2]
			r += middle
		}
	}
	return r, nil
}

func parttwo(lines []string) (r int, err error) {
	pages, updates := read_data(lines)

	for _, update := range updates {
		if ok, badPage := validate_update(pages, update); !ok {
			//fmt.Printf("Fail for %v, bad page idx is %v\n", update, badPage)
			middle := fix_that_update(pages, update, badPage)
			//fmt.Printf("I totally fixed it though: %v\n", r)
			r += middle
		}
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
