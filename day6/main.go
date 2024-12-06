package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func parseInput() [][]string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	area := make([][]string, 0)

	for _, line := range strings.Split(string(input), "\n") {
		area = append(area, strings.Split(line, ""))
	}

	return area
}

func getNextPosition(x, y int, currentPos string, area [][]string) (int, int, string, [][]string, bool) {
	area[x][y] = "."
	switch currentPos {
	case "^":
		x--
		if x == -1 {
			return x, y, currentPos, area, true
		}
		if area[x][y] == "." {
			area[x][y] = "^"
			return x, y, "^", area, false
		}
		area[x+1][y] = ">"
		return x + 1, y, ">", area, false
	case "v":
		x++
		if x == len(area) {
			return x, y, currentPos, area, true
		}
		if area[x][y] == "." {
			area[x][y] = "v"
			return x, y, "v", area, false
		}
		area[x-1][y] = "<"
		return x - 1, y, "<", area, false
	case "<":
		y--
		if y == -1 {
			return x, y, currentPos, area, true
		}
		if area[x][y] == "." {
			area[x][y] = "<"
			return x, y, "<", area, false
		}
		area[x][y+1] = "^"
		return x, y + 1, "^", area, false
	case ">":
		y++
		if y == len(area[0]) {
			return x, y, currentPos, area, true
		}
		if area[x][y] == "." {
			area[x][y] = ">"
			return x, y, ">", area, false
		}
		area[x][y-1] = "v"
		return x, y - 1, "v", area, false
	}

	panic("Invalid position")
}

func findAWayOut(area [][]string) (bool, map[[2]int]bool) {
	startX, startY := 0, 0
	for x, row := range area {
		if y := slices.Index(row, "^"); y != -1 {
			startX, startY = x, y
			break
		}
	}

	x, y := 0, 0
	pos := "^"
	isOut := false

	visited := make(map[[2]int]bool)
	visited[[2]int{startX, startY}] = true
	x, y = startX, startY

	size := len(visited)
	movesWithoutChange := 0

	for !isOut {
		x, y, pos, area, isOut = getNextPosition(x, y, pos, area)

		if !isOut {
			visited[[2]int{x, y}] = true

			if size != len(visited) {
				movesWithoutChange = 0
				size = len(visited)
			} else if movesWithoutChange == 200 {
				// We are stuck in a loop
				break
			} else {
				movesWithoutChange++
			}
		}
	}

	return isOut, visited
}

func part1() {
	area := parseInput()

	_, visited := findAWayOut(area)

	fmt.Printf("Result part 1: %d\n", len(visited))
}

func part2() {
	area := parseInput()
	_, visited := findAWayOut(area)

	res := 0

	for k := range visited {
		area = parseInput()
		area[k[0]][k[1]] = "#"

		isOut, _ := findAWayOut(area)
		if !isOut {
			res++
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
