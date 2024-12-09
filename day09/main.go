package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type Block struct {
	ID   int
	Size int
	Prev *Block
	Next *Block
}

func readData(line string) (first *Block, last *Block) {

	id := 0
	var current *Block
	writingFile := true
	for _, v := range line {
		iv, _ := strconv.Atoi(string(v))
		for i := 0; i < iv; i++ {
			if writingFile {
				newblock := &Block{
					ID:   id,
					Size: iv,
					Prev: current,
				}
				if first == nil {
					first = newblock
					current = newblock
				} else {
					current.Next = newblock
					current = newblock
				}
			} else {
				newblock := &Block{
					ID:   -1,
					Size: iv,
					Prev: current,
				}
				current.Next = newblock
				current = newblock
			}
		}
		if writingFile {
			id += 1
		}
		writingFile = !writingFile
	}
	last = current
	return first, last
}

func getNextFree(b *Block) *Block {
	var nextFree *Block = b
	for nextFree != nil && nextFree.ID != -1 {
		nextFree = nextFree.Next
	}
	return nextFree
}

func getNextFreeOfSize(b *Block, fileToMove *Block) *Block {
	var nextFree *Block = b
	for nextFree != nil && (nextFree.ID != -1 || (nextFree.ID == -1 && nextFree.Size < fileToMove.Size)) {
		nextFree = nextFree.Next
		if nextFree == fileToMove {
			// We shouldn't move the file after where it currently sits, so we failed to find
			// free space
			nextFree = nil
			break
		}
	}
	return nextFree
}

func getPrevFile(b *Block) *Block {
	var prevFile *Block = b
	for prevFile != nil && prevFile.ID == -1 {
		prevFile = prevFile.Prev
	}
	return prevFile
}

func getPrevFileOfId(b *Block, id int) *Block {
	var prevFile *Block = b
	for prevFile != nil && prevFile.ID != id {
		prevFile = prevFile.Prev
	}
	return prevFile
}

func compress(first *Block, last *Block) {
	nfree := getNextFree(first)
	flast := getPrevFile(last)

	for {
		nfree.ID = flast.ID
		flast.ID = -1
		nfree = getNextFree(nfree)
		flast = getPrevFile(last)

		if nfree.Prev == flast {
			break
		}
	}
}

func compress2(first *Block, last *Block) {
	flast := getPrevFile(last)

	nextFileId := flast.ID

	for nextFileId >= 0 && flast != nil {
		//fmt.Printf("Looking for free space for ID %v of size %v\n", flast.ID, flast.Size)
		nextFree := getNextFreeOfSize(first, flast)
		if nextFree != nil {
			// There is space available

			// We need to update the remaining free blocks in this slice to
			// their new free contiguous space size after all this
			remainingFreeSize := nextFree.Size - flast.Size

			//fmt.Printf("Moving ID %v of size %v\n", flast.ID, flast.Size)
			sizeToMove := flast.Size
			for i := 0; i < sizeToMove; i++ {
				nextFree.ID = flast.ID
				nextFree = nextFree.Next
				flast.ID = -1
				flast = flast.Prev
			}

			// Update the remaining size of the free blocks
			//fmt.Printf("Setting free remaining size to %v\n", remainingFreeSize)
			for nextFree.ID == -1 {
				nextFree.Size = remainingFreeSize
				nextFree = nextFree.Next
			}
		}
		nextFileId -= 1
		flast = getPrevFileOfId(flast, nextFileId)
		//printBlock(first)
	}
}

func sum(first *Block) (r int) {
	var bidx = first
	position := 0
	for bidx != nil {
		if bidx.ID >= 0 {
			r += position * bidx.ID
		}
		bidx = bidx.Next
		position++
	}
	return r
}

func partone(line string) (r int, err error) {
	first, last := readData(line)
	compress(first, last)
	r = sum(first)
	return r, nil
}

func parttwo(line string) (r int, err error) {
	first, last := readData(line)
	compress2(first, last)
	r = sum(first)
	return r, nil
}

func printBlock(b *Block) {
	n := b
	for n != nil {
		if n.ID >= 0 {
			fmt.Printf("%v", n.ID)
		} else {
			fmt.Printf(".")
		}
		n = n.Next
	}
	fmt.Printf("\n")
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partone(lines[0])
	fmt.Printf("Part one: %v\n", ans)

	ans, err = parttwo(lines[0])
	fmt.Printf("Part two: %v\n", ans)

}
