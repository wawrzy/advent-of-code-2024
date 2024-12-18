package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	gridSize = 71
)

func parseInput(maxBytes int) ([][]string, []string) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	grid := make([][]string, 0)

	for i := 0; i < gridSize; i++ {
		row := make([]string, 0)
		for j := 0; j < gridSize; j++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	for i, line := range lines {
		if i == maxBytes {
			break
		}

		b := strings.Split(line, ",")
		x, _ := strconv.Atoi(b[0])
		y, _ := strconv.Atoi(b[1])

		grid[y][x] = "#"
	}

	return grid, lines
}

type Node struct {
	G, H, F int
	IsWall  bool
}

type AStar struct {
	Grid [][]*Node

	StartX, StartY, EndX, EndY int

	closeList [][2]int
	openList  [][2]int
}

func NewAStar(grid [][]string, startX, startY, endX, endY int) *AStar {
	nodeGrid := make([][]*Node, 0)

	for _, row := range grid {
		nodeRow := make([]*Node, 0)
		for _, cell := range row {
			nodeRow = append(nodeRow, &Node{
				G:      0,
				H:      0,
				F:      0,
				IsWall: cell == "#",
			})
		}
		nodeGrid = append(nodeGrid, nodeRow)
	}

	aStar := &AStar{
		Grid: nodeGrid,

		StartX: startX,
		StartY: startY,
		EndX:   endX,
		EndY:   endY,

		closeList: make([][2]int, 0),
		openList:  make([][2]int, 0),
	}

	aStar.openList = append(aStar.openList, [2]int{startX, startY})

	return aStar
}

func (a *AStar) getManhattanDistance(x, y int) int {
	return int(math.Abs(float64(a.EndX-x)) + math.Abs(float64(a.EndY-y)))
}

func (a *AStar) getNeighbors(x, y int) [][2]int {
	neighbors := make([][2]int, 0)
	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for _, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]

		if newX >= 0 && newX < gridSize && newY >= 0 && newY < gridSize {
			if !a.Grid[newY][newX].IsWall {
				neighbors = append(neighbors, [2]int{newX, newY})
			}
		}
	}

	return neighbors
}

func (a *AStar) popLowestF() [2]int {
	lowestNodeF := a.openList[0]
	idx := 0

	for i, pos := range a.openList {
		if i == 0 {
			continue
		}

		if a.Grid[pos[1]][pos[0]].F < a.Grid[lowestNodeF[1]][lowestNodeF[0]].F {
			lowestNodeF = pos
			idx = i
		}
	}

	a.openList = append(a.openList[:idx], a.openList[idx+1:]...)
	a.closeList = append(a.closeList, lowestNodeF)

	return lowestNodeF
}

func (a *AStar) findShortestPath() int {
	for len(a.openList) > 0 {
		current := a.popLowestF()

		if current[0] == a.EndX && current[1] == a.EndY {
			break
		}

		neighbors := a.getNeighbors(current[0], current[1])
		for _, neighbor := range neighbors {
			if slices.Contains(a.closeList, neighbor) || slices.Contains(a.openList, neighbor) {
				continue
			}

			neighborNode := a.Grid[neighbor[1]][neighbor[0]]

			neighborNode.G = a.Grid[current[1]][current[0]].G + 1
			neighborNode.H = a.getManhattanDistance(neighbor[0], neighbor[1])
			neighborNode.F = neighborNode.G + neighborNode.H

			a.openList = append(a.openList, neighbor)
		}
	}

	return a.Grid[a.EndY][a.EndX].F
}

func part1() {
	grid, _ := parseInput(1024)

	astar := NewAStar(grid, 0, 0, gridSize-1, gridSize-1)

	res := astar.findShortestPath()

	fmt.Printf("Result part 1: %d\n", res)
}

func part2() {
	_, bytes := parseInput(0)

	minMaxBytes := 1024
	maxMaxBytes := len(bytes)

	for {
		currentMaxBytes := (minMaxBytes + maxMaxBytes) / 2

		grid, _ := parseInput(currentMaxBytes)

		astar := NewAStar(grid, 0, 0, gridSize-1, gridSize-1)

		res := astar.findShortestPath()

		if res == 0 {
			maxMaxBytes = currentMaxBytes
		} else {
			minMaxBytes = currentMaxBytes
		}

		if minMaxBytes+1 == maxMaxBytes {
			break
		}
	}

	fmt.Printf("Result part 2: %s\n", bytes[maxMaxBytes-1])
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
