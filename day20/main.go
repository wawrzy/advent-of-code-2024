package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"advent/astar"
)

const shouldSave = 100

func parseInput() [][]string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	grid := make([][]string, 0)

	for _, line := range lines {
		row := make([]string, 0)

		for _, char := range line {
			row = append(row, string(char))
		}

		grid = append(grid, row)
	}

	return grid
}

func getManahttanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

func countPossiblePaths(cheatAllowed int) int {
	grid := parseInput()

	star := astar.NewAStar(grid)

	_, path := star.FindShortestPath()

	res := 0

	for i := 0; i < len(path); i++ {
		steps := i + shouldSave + 1
		orig := path[i]

		for steps < len(path) {
			end := path[steps]

			manathanDistance := getManahttanDistance(orig[0], orig[1], end[0], end[1])

			if manathanDistance <= cheatAllowed && steps-i-manathanDistance >= shouldSave {
				res++
			}

			steps++
		}
	}

	return res
}

func part1() {
	fmt.Printf("Result part 1: %d\n", countPossiblePaths(2))
}

func part2() {
	fmt.Printf("Result part 2: %d\n", countPossiblePaths(20))
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
