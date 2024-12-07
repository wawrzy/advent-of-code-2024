package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Equation struct {
	Total     int
	Parts     []int
	Operators []string
	IsValid   bool
}

func parseInput() []Equation {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")

	equations := make([]Equation, 0, len(lines))

	for _, line := range lines {
		equation := Equation{}

		mainParts := strings.Split(line, ": ")
		if len(mainParts) != 2 {
			break
		}

		total, err := strconv.Atoi(mainParts[0])
		if err != nil {
			panic(err)
		}
		equation.Total = total

		parts := strings.Split(mainParts[1], " ")
		equation.Parts = make([]int, 0, len(parts))

		for _, part := range parts {
			partValue, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			equation.Parts = append(equation.Parts, partValue)
		}

		equations = append(equations, equation)
	}

	return equations
}

func generateEquationOperators(size int) map[int][][]string {
	operators := make(map[int][][]string, size)

	for currentSize := 1; currentSize <= size; currentSize++ {
		result := make([][]string, 0, 2^size)

		firstRow := make([]string, 0, size)
		for i := 0; i < size; i++ {
			firstRow = append(firstRow, "+")
		}

		result = append(result, firstRow)

		for i := 0; i < size; i++ {
			newRows := make([][]string, 0, len(result))

			for _, row := range result {
				newRow := make([]string, 0, size)
				newRow = append(newRow, row...)
				newRow[i] = "*"
				newRows = append(newRows, newRow)
			}

			result = append(result, newRows...)
		}

		operators[currentSize] = result
	}

	return operators
}

func solveEquation(equation Equation, operators []string) int {
	result := equation.Parts[0]
	for i, part := range equation.Parts[1:] {
		operator := operators[i]
		if operator == "+" {
			result += part
		} else if operator == "*" {
			result *= part
		}
	}

	return result
}

func isValidEquation(equation Equation, solutions [][]string) bool {
	for _, solution := range solutions {
		if solveEquation(equation, solution) == equation.Total {
			return true
		}
	}
	return false
}

func part1() {
	equations := parseInput()

	maxParts := 0
	for _, equation := range equations {
		if len(equation.Parts) > maxParts {
			maxParts = len(equation.Parts)
		}
	}

	operators := generateEquationOperators(maxParts - 1)

	res := 0

	for _, equation := range equations {
		isValid := isValidEquation(equation, operators[len(equation.Parts)-1])

		if isValid {
			res += equation.Total
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

type OperatorsForSize struct {
	MultOperators   [][]string
	ConcatOperators [][]string
	BothOperators   [][]string
}

func generateEquationOperatorsPart2(size int) map[int]OperatorsForSize {
	operators := make(map[int]OperatorsForSize, size)

	for currentSize := 1; currentSize <= size; currentSize++ {
		var wg sync.WaitGroup

		resultMult := make([][]string, 0, 3^currentSize)
		resultConcat := make([][]string, 0, 3^currentSize)
		resultBoth := make([][]string, 0, 3^currentSize)

		wg.Add(1)
		go func() {
			defer wg.Done()

			firstRow := make([]string, 0, currentSize)
			for i := 0; i < currentSize; i++ {
				firstRow = append(firstRow, "+")
			}

			resultMult = append(resultMult, firstRow)

			for i := 0; i < currentSize; i++ {
				newRows := make([][]string, 0, len(resultMult))

				for _, row := range resultMult {
					newRow := make([]string, 0, currentSize)
					newRow = append(newRow, row...)
					newRow[i] = "*"
					newRows = append(newRows, newRow)
				}

				resultMult = append(resultMult, newRows...)
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			firstRow := make([]string, 0, currentSize)
			for i := 0; i < currentSize; i++ {
				firstRow = append(firstRow, "+")
			}

			resultConcat = append(resultConcat, firstRow)

			for i := 0; i < currentSize; i++ {
				newRows := make([][]string, 0, len(resultConcat))

				for _, row := range resultConcat {
					newRow := make([]string, 0, currentSize)
					newRow = append(newRow, row...)
					newRow[i] = "||"
					newRows = append(newRows, newRow)
				}

				resultConcat = append(resultConcat, newRows...)
			}
		}()

		wg.Wait()

		firstRows := make([][]string, 0, len(resultMult))
		for _, row := range resultMult {
			newRow := make([]string, 0, currentSize)
			newRow = append(newRow, row...)
			firstRows = append(firstRows, newRow)
		}

		resultBoth = append(resultBoth, firstRows...)

		for i := 0; i < currentSize; i++ {
			newRows := make([][]string, 0, len(resultBoth))

			for _, row := range resultBoth {
				newRow := make([]string, 0, currentSize)
				newRow = append(newRow, row...)
				newRow[i] = "||"
				newRows = append(newRows, newRow)
			}

			resultBoth = append(resultBoth, newRows...)
		}

		operators[currentSize] = OperatorsForSize{
			MultOperators:   resultMult,
			ConcatOperators: resultConcat,
			BothOperators:   resultBoth,
		}
	}

	return operators
}

func solveEquationPart2(equation Equation, operators []string) int {
	result := equation.Parts[0]
	for i, part := range equation.Parts[1:] {
		operator := operators[i]
		if operator == "+" {
			result += part
		} else if operator == "*" {
			result *= part
		} else if operator == "||" {
			resultStr := strconv.Itoa(result)
			partStr := strconv.Itoa(part)
			result, _ = strconv.Atoi(resultStr + partStr)
		}
	}

	return result
}

func isValidEquationPart2(equation Equation, solutions OperatorsForSize) bool {
	for _, solution := range solutions.MultOperators {
		if solveEquationPart2(equation, solution) == equation.Total {
			return true
		}
	}

	for _, solution := range solutions.ConcatOperators {
		if solveEquationPart2(equation, solution) == equation.Total {
			return true
		}
	}

	for _, solution := range solutions.BothOperators {
		if solveEquationPart2(equation, solution) == equation.Total {
			return true
		}
	}
	return false
}

func part2() {
	equations := parseInput()

	maxParts := 0
	for _, equation := range equations {
		if len(equation.Parts) > maxParts {
			maxParts = len(equation.Parts)
		}
	}

	operators := generateEquationOperatorsPart2(maxParts - 1)

	res := 0

	for _, equation := range equations {
		isValid := isValidEquationPart2(equation, operators[len(equation.Parts)-1])

		if isValid {
			res += equation.Total
		}
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
