package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseInput() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(input)
}

func findNextExpression(program string, expression string) int {
	return strings.Index(program, expression)
}

func part1() {
	program := parseInput()

	res := 0

	for idx := findNextExpression(program, "mul("); idx != -1; idx = findNextExpression(program, "mul(") {
		expressionEnd := findNextExpression(program[idx:], ")") + idx
		expression := program[idx : expressionEnd+1]

		values := strings.Split(expression[4:len(expression)-1], ",")

		if len(values) == 2 && len(values[0]) <= 3 && len(values[1]) <= 3 {
			left, errL := strconv.Atoi(values[0])
			right, errR := strconv.Atoi(values[1])

			if errL == nil && errR == nil {
				res += left * right
			}

		}

		program = program[idx+4:]
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func part2() {
	program := parseInput()

	res := 0
	enabled := true

	idx := 0
	for idx < len(program) {
		nextDo := findNextExpression(program[idx:], "do()")
		nextDont := findNextExpression(program[idx:], "don't()")
		nextMul := findNextExpression(program[idx:], "mul(")

		if nextMul == -1 {
			break
		}

		if nextDo != -1 && (nextDont == -1 || nextDo < nextDont) && nextDo < nextMul {
			enabled = true
			idx += nextDo + 1
			continue
		} else if nextDont != -1 && (nextDo == -1 || nextDont < nextDo) && nextDont < nextMul {
			enabled = false
			idx += nextDont + 1
			continue
		}

		idx += nextMul

		if !enabled {
			idx += 1
			continue
		}

		expressionEnd := findNextExpression(program[idx:], ")") + idx
		expression := program[idx : expressionEnd+1]

		values := strings.Split(expression[4:len(expression)-1], ",")

		if len(values) == 2 && len(values[0]) <= 3 && len(values[1]) <= 3 {
			left, errL := strconv.Atoi(values[0])
			right, errR := strconv.Atoi(values[1])

			if errL == nil && errR == nil {
				res += left * right
			}
		}

		idx += 1
	}

	fmt.Printf("Result part 2: %d\n", res)
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
