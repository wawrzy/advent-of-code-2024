package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func reverseArray(row []string) []string {
	newRow := make([]string, len(row))

	for i := 0; i <= len(row)/2; i++ {
		j := len(row) - i - 1
		newRow[i], newRow[j] = row[j], row[i]
	}

	return newRow
}

func parseInput() [][]string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	result := make([][]string, 0, len(lines))

	for _, line := range lines {
		result = append(result, strings.Split(line, ""))
	}

	return result
}

func searchHorizontal(grid [][]string, reverse bool) int {
	res := 0

	for _, row := range grid {
		rowString := ""

		if reverse {
			rowString = strings.Join(reverseArray(row), "")
		} else {
			rowString = strings.Join(row, "")
		}

		for idx := strings.Index(rowString, "XMAS"); idx != -1; idx = strings.Index(rowString, "XMAS") {
			res++
			rowString = rowString[idx+1:]
		}
	}

	return res
}

func searchVertical(grid [][]string, reverse bool) int {
	res := 0

	for i := 0; i < len(grid[0]); i++ {
		col := make([]string, 0, len(grid))

		for j := 0; j < len(grid); j++ {
			col = append(col, grid[j][i])
		}

		colString := ""

		if reverse {
			colString = strings.Join(reverseArray(col), "")
		} else {
			colString = strings.Join(col, "")
		}

		for idx := strings.Index(colString, "XMAS"); idx != -1; idx = strings.Index(colString, "XMAS") {
			res++
			colString = colString[idx+1:]
		}
	}

	return res
}

func searchDiagonal(grid [][]string, reverse bool) int {
	res := 0

	rows := make([]string, 0, len(grid)*2-1)

	for i := 0; i < len(grid)*2-1; i++ {
		x := i

		if reverse {
			if i >= len(grid) {
				x = len(grid) - 1
			}
		} else if i >= len(grid) {
			x = len(grid) - 1
		}

		y := 0

		if reverse {
			y = len(grid) - 1

			if i >= len(grid) {
				y = len(grid) - 2 - i%len(grid)
			}
		} else if i >= len(grid) {
			y = i%len(grid) + 1
		}

		diagonal := make([]string, 0, len(grid))

		for x >= 0 && y < len(grid[0]) && y >= 0 {
			diagonal = append(diagonal, grid[x][y])

			if reverse {
				x--
				y--
			} else {
				x--
				y++
			}
		}

		rows = append(rows, strings.Join(diagonal, ""))
	}

	for _, row := range rows {
		inversedRow := strings.Join(reverseArray(strings.Split(row, "")), "")

		for idx := strings.Index(inversedRow, "XMAS"); idx != -1; idx = strings.Index(inversedRow, "XMAS") {
			res++
			inversedRow = inversedRow[idx+1:]
		}
	}

	for _, row := range rows {
		for idx := strings.Index(row, "XMAS"); idx != -1; idx = strings.Index(row, "XMAS") {
			res++
			row = row[idx+1:]
		}
	}

	return res
}

func part1() {
	input := parseInput()

	res := 0

	res += searchHorizontal(input, false)
	res += searchHorizontal(input, true)
	res += searchVertical(input, false)
	res += searchVertical(input, true)
	res += searchDiagonal(input, false)
	res += searchDiagonal(input, true)

	fmt.Printf("Result part 1: %d\n", res)
}

func isXMAS(grid [][]string, x, y int) bool {
	if x == 0 || y == len(grid[0])-1 || x == len(grid)-1 || y == 0 {
		return false
	}

	n := 0

	// top left
	if grid[x-1][y-1] == "M" {
		n++
		// bottom right
		if grid[x+1][y+1] != "S" {
			return false
		}
	}

	// top right
	if grid[x-1][y+1] == "M" {
		n++
		// bottom left
		if grid[x+1][y-1] != "S" {
			return false
		}
	}

	// bottom right
	if grid[x+1][y+1] == "M" {
		n++
		// top left
		if grid[x-1][y-1] != "S" {
			return false
		}
	}

	// bottom left
	if grid[x+1][y-1] == "M" {
		n++
		// top right
		if grid[x-1][y+1] != "S" {
			return false
		}
	}

	return n == 2
}

func part2() {
	input := parseInput()

	res := 0

	for x, row := range input {
		for y, cell := range row {
			if cell == "A" {
				if isXMAS(input, x, y) {
					res++
				}
			}
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
