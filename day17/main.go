package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2024-golang/helper"
)

type OpCode int

const (
	ADV OpCode = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

type CPU struct {
	A       int
	B       int
	C       int
	I       int
	Program []int
	Output  []int
}

func readData(lines []string) CPU {
	cpu := CPU{
		A:       0,
		B:       0,
		C:       0,
		I:       0,
		Program: make([]int, 0),
		Output:  make([]int, 0),
	}
	cpu.A, _ = strconv.Atoi(strings.Split(lines[0], " ")[2])
	cpu.B, _ = strconv.Atoi(strings.Split(lines[1], " ")[2])
	cpu.C, _ = strconv.Atoi(strings.Split(lines[2], " ")[2])
	regstr := strings.Split(strings.Split(lines[3], " ")[1], ",")
	for _, r := range regstr {
		ri, _ := strconv.Atoi(r)
		cpu.Program = append(cpu.Program, ri)
	}
	return cpu
}

func (c *CPU) Instruction(seen map[string]struct{}) bool {
	if c.I >= len(c.Program)-1 {
		return false
	}
	if _, ok := seen[c.seen()]; ok {
		// been in this state before and we know it's not good
		return false
	}
	seen[c.seen()] = struct{}{}

	operand := c.Program[c.I+1]

	// special flag if we're doing JNZ as it impacts post-instruction logic
	jumped := false
	switch OpCode(c.Program[c.I]) {
	case ADV:
		c.adv(operand)
	case BXL:
		c.bxl(operand)
	case BST:
		c.bst(operand)
	case JNZ:
		jumped = c.jnz(operand)
	case BXC:
		c.bxc(operand)
	case OUT:
		c.out(operand)
	case BDV:
		c.bdv(operand)
	case CDV:
		c.cdv(operand)
	}

	if !jumped {
		c.I += 2
	}

	return true
}

func (c *CPU) adv(o int) {
	n := c.A
	d := int(math.Pow(float64(2), float64(c.combo(o))))
	c.A = n / d
}

func (c *CPU) bxl(o int) {
	c.B = c.B ^ o
}

func (c *CPU) bst(o int) {
	c.B = c.combo(o) % 8
}

func (c *CPU) jnz(o int) bool {
	if c.A == 0 {
		return false
	}
	c.I = o
	return true
}

func (c *CPU) bxc(o int) {
	c.B = c.B ^ c.C
}

func (c *CPU) out(o int) {
	c.Output = append(c.Output, c.combo(o)%8)
}

func (c *CPU) bdv(o int) {
	n := c.A
	d := int(math.Pow(float64(2), float64(c.combo(o))))
	c.B = n / d
}

func (c *CPU) cdv(o int) {
	n := c.A
	d := int(math.Pow(float64(2), float64(c.combo(o))))
	c.C = n / d
}

func (c *CPU) seen() string {
	return fmt.Sprintf("%v%v%v%v", c.A, c.B, c.C, c.I)
}

func (c *CPU) combo(o int) int {
	if o >= 0 && o <= 3 {
		return o
	} else if o == 4 {
		return c.A
	} else if o == 5 {
		return c.B
	} else if o == 6 {
		return c.C
	} else {
		return -1
	}
}

func (c *CPU) PrintOutput() string {
	soutput := make([]string, len(c.Output))
	for i, num := range c.Output {
		soutput[i] = strconv.Itoa(num)
	}
	return strings.Join(soutput, ",")
}

func (c *CPU) PrintProgram() string {
	soutput := make([]string, len(c.Program))
	for i, num := range c.Program {
		soutput[i] = strconv.Itoa(num)
	}
	return strings.Join(soutput, ",")
}

func (c *CPU) PrintState() {
	fmt.Printf("A: %v, B: %b, C: %v, I: %v, O: %v\n", c.A, c.B, c.C, c.I, c.Output)
}

func (c *CPU) Clone() CPU {
	r := CPU{
		A:       c.A,
		B:       c.B,
		C:       c.C,
		I:       c.I,
		Program: make([]int, 0),
		Output:  make([]int, 0),
	}
	for _, v := range c.Program {
		r.Program = append(r.Program, v)
	}
	return r
}

func partone(lines []string) (r int, err error) {
	cpu := readData(lines)
	seen := make(map[string]struct{})
	for cpu.Instruction(seen) {
		//cpu.PrintState()
	}
	fmt.Printf("%v\n", cpu.PrintOutput())
	return 0, nil
}

func parttwo(lines []string) (r int, err error) {
	cpu := readData(lines)

	testOutput := cpu.PrintProgram()
	acounter := 1
	seen := make(map[string]struct{})
	for {
		trial := cpu.Clone()
		trial.A = acounter

		for trial.Instruction(seen) {
		}

		if trial.PrintOutput() == testOutput {
			break
		}
		acounter++
	}
	return acounter, nil
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
