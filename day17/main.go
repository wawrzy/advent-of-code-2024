package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Computer struct {
	RegisterA int
	RegisterB int
	RegisterC int

	Pointer int
	Output  string

	Program []int
}

func (c *Computer) comboOperand() int {
	operand := c.Program[c.Pointer+1]

	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.RegisterA
	case 5:
		return c.RegisterB
	case 6:
		return c.RegisterC
	}

	panic("Invalid operand")
}

func (c *Computer) literalOperand() int {
	return c.Program[c.Pointer+1]
}

func (c *Computer) adv() {
	numerator := c.RegisterA
	denominator := int(math.Pow(2, float64(c.comboOperand())))

	c.RegisterA = numerator / denominator
}

func (c *Computer) bxl() {
	c.RegisterB ^= c.literalOperand()
}

func (c *Computer) bst() {
	c.RegisterB = c.comboOperand() % 8
}

func (c *Computer) jnz() {
	if c.RegisterA == 0 {
		return
	}
	c.Pointer = c.literalOperand() - 2
}

func (c *Computer) bxc() {
	c.RegisterB = c.RegisterB ^ c.RegisterC
}

func (c *Computer) out() {
	value := c.comboOperand() % 8
	str := strconv.Itoa(value)

	str = strings.Join(strings.Split(str, ""), ",")

	if c.Output == "" {
		c.Output = str
	} else {
		c.Output = fmt.Sprintf("%s,%s", c.Output, str)
	}
}

func (c *Computer) bdv() {
	numerator := c.RegisterA
	denominator := int(math.Pow(2, float64(c.comboOperand())))

	c.RegisterB = numerator / denominator
}

func (c *Computer) cdv() {
	numerator := c.RegisterA
	denominator := int(math.Pow(2, float64(c.comboOperand())))

	c.RegisterC = numerator / denominator
}

func (c *Computer) reset() {
	c.RegisterA = 0
	c.RegisterB = 0
	c.RegisterC = 0
	c.Pointer = 0
	c.Output = ""
}

func (c *Computer) execute() {
	for c.Pointer < len(c.Program) {
		opcode := c.Program[c.Pointer]
		switch opcode {
		case 0:
			c.adv()
		case 1:
			c.bxl()
		case 2:
			c.bst()
		case 3:
			c.jnz()
		case 4:
			c.bxc()
		case 5:
			c.out()
		case 6:
			c.bdv()
		case 7:
			c.cdv()
		}
		c.Pointer += 2
	}
}

func (c *Computer) OutputToNumbers() []int {
	numbers := make([]int, 0)

	for _, c := range strings.Split(c.Output, ",") {
		value, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, value)
	}

	return numbers
}

func parseInput() (Computer, string) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	registerAStr := strings.Split(lines[0], " ")[2]
	registerBStr := strings.Split(lines[1], " ")[2]
	registerCStr := strings.Split(lines[2], " ")[2]

	registerA, err := strconv.Atoi(registerAStr)
	if err != nil {
		panic(err)
	}

	registerB, err := strconv.Atoi(registerBStr)
	if err != nil {
		panic(err)
	}

	registerC, err := strconv.Atoi(registerCStr)
	if err != nil {
		panic(err)
	}

	program := make([]int, 0)

	programStr := strings.Split(lines[4], " ")[1]
	for _, c := range strings.Split(programStr, ",") {
		value, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		program = append(program, value)
	}

	return Computer{
		RegisterA: registerA,
		RegisterB: registerB,
		RegisterC: registerC,
		Program:   program,
	}, programStr
}

func part1() {
	computer, _ := parseInput()

	computer.execute()

	fmt.Printf("Result part 1: %s\n", computer.Output)
}

func part2() {
	computer, expect := parseInput()

	registerA := 1
	for i := 0; i < len(computer.Program)-1; i++ {
		registerA += 7 * int(math.Pow(8, float64(i)))
	}

	for {
		computer.reset()
		computer.RegisterA = registerA
		computer.execute()

		if computer.Output == expect {
			break
		}

		output := computer.OutputToNumbers()
		add := 0
		for i := len(output) - 1; i >= 0; i-- {
			if output[i] != computer.Program[i] {
				add = int(math.Pow(8, float64(i)))
				registerA += add
				break
			}
		}
	}

	fmt.Printf("Result part 2: %d\n", registerA)
}

func main() {
	now := time.Now()
	part1()
	duration := time.Since(now)

	fmt.Printf("Duration part 1: %v\n", duration)

	now = time.Now()
	part2()
	duration = time.Since(now)

	fmt.Printf("Duration part 2: %v\n", duration)
}
