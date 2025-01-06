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
	A       uint64
	B       uint64
	C       uint64
	I       uint64
	Program []uint64
	Output  []uint64
}

func readData(lines []string) CPU {
	cpu := CPU{
		A:       0,
		B:       0,
		C:       0,
		I:       0,
		Program: make([]uint64, 0),
		Output:  make([]uint64, 0),
	}
	a, _ := strconv.Atoi(strings.Split(lines[0], " ")[2])
	b, _ := strconv.Atoi(strings.Split(lines[1], " ")[2])
	c, _ := strconv.Atoi(strings.Split(lines[2], " ")[2])
	cpu.A = uint64(a)
	cpu.B = uint64(b)
	cpu.C = uint64(c)
	regstr := strings.Split(strings.Split(lines[3], " ")[1], ",")
	for _, r := range regstr {
		ri, _ := strconv.Atoi(r)
		cpu.Program = append(cpu.Program, uint64(ri))
	}
	return cpu
}

func (c *CPU) Instruction() bool {
	if c.I >= uint64(len(c.Program)-1) {
		return false
	}

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

func (c *CPU) adv(o uint64) {
	n := c.A
	d := uint64(math.Pow(float64(2), float64(c.combo(o))))
	c.A = n / d
}

func (c *CPU) bxl(o uint64) {
	c.B = c.B ^ o
}

func (c *CPU) bst(o uint64) {
	c.B = c.combo(o) % 8
}

func (c *CPU) jnz(o uint64) bool {
	if c.A == 0 {
		return false
	}
	c.I = o
	return true
}

func (c *CPU) bxc(o uint64) {
	c.B = c.B ^ c.C
}

func (c *CPU) out(o uint64) {
	c.Output = append(c.Output, c.combo(o)%8)
}

func (c *CPU) bdv(o uint64) {
	n := c.A
	d := uint64(math.Pow(float64(2), float64(c.combo(o))))
	c.B = n / d
}

func (c *CPU) cdv(o uint64) {
	n := c.A
	d := uint64(math.Pow(float64(2), float64(c.combo(o))))
	c.C = n / d
}

func (c *CPU) combo(o uint64) uint64 {
	if o >= 0 && o <= 3 {
		return o
	} else if o == 4 {
		return c.A
	} else if o == 5 {
		return c.B
	} else if o == 6 {
		return c.C
	} else {
		return 0
	}
}

func (c *CPU) PrintOutput() string {
	soutput := make([]string, len(c.Output))
	for i, num := range c.Output {
		soutput[i] = strconv.FormatUint(num, 10)
	}
	return strings.Join(soutput, ",")
}

func (c *CPU) PrintProgram() string {
	soutput := make([]string, len(c.Program))
	for i, num := range c.Program {
		soutput[i] = strconv.FormatUint(num, 10)
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
		Program: make([]uint64, 0),
		Output:  make([]uint64, 0),
	}
	for _, v := range c.Program {
		r.Program = append(r.Program, v)
	}
	return r
}

func logN(a uint64, n uint64) uint64 {
	result := math.Log(float64(a)) / math.Log(float64(n))
	return uint64(result)
}

func partone(lines []string) (r uint64, err error) {
	cpu := readData(lines)
	for cpu.Instruction() {
		//cpu.PrintState()
	}
	fmt.Printf("%v\n", cpu.PrintOutput())
	return 0, nil
}

func parttwo(lines []string) (r uint64, err error) {
	cpu := readData(lines)

	pow := uint64(0)
	aVal := uint64(math.Pow(8, float64(len(cpu.Program)-1)))

	for {
		cpu.A = aVal
		for cpu.Instruction() {
		}
		if cpu.PrintOutput() == cpu.PrintProgram() {
			break
		}

		pow = logN(aVal, 8)
		for i := 0; i < len(cpu.Output)-1; i++ {
			n := cpu.Output[len(cpu.Output)-1-i]
			programIndex := len(cpu.Program) - 1 - i

			if programIndex < 0 || cpu.Program[programIndex] != n {
				break
			}
			pow--
		}

		aVal += uint64(math.Pow(8, float64(pow)))
		cpu.B = 0
		cpu.C = 0
		cpu.I = 0
		cpu.Output = make([]uint64, 0)
	}
	return aVal, nil
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
