package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

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

func NewAStar(grid [][]string) *AStar {
	nodeGrid := make([][]*Node, 0)

	startX, startY := 0, 0
	endX, endY := 0, 0

	for y, row := range grid {
		nodeRow := make([]*Node, 0)
		for x, cell := range row {
			nodeRow = append(nodeRow, &Node{
				G:      0,
				H:      0,
				F:      0,
				IsWall: cell == "#",
			})

			if cell == "S" {
				startX = x
				startY = y
			} else if cell == "E" {
				endX = x
				endY = y
			}
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

		if newX >= 0 && newX < len(a.Grid) && newY >= 0 && newY < len(a.Grid[0]) {
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
	grid := parseInput()

	astar := NewAStar(grid)
	baseTime := astar.findShortestPath()

	var mutex sync.Mutex
	var waitGroup sync.WaitGroup

	timeSaved := make(map[int]int)

	y := 0

	for y+10 < len(grid) {
		waitGroup.Add(1)
		go func(y int) {
			defer waitGroup.Done()

			grid := parseInput()

			maxY := y + 10
			for y := y; y < maxY && y < len(grid); y += 1 {
				for x := 0; x < len(grid[y]); x += 1 {
					if x == 0 || x == len(grid[y])-1 {
						continue
					}

					if grid[y][x] == "#" {
						grid[y][x] = "."
						astar = NewAStar(grid)
						time := astar.findShortestPath()

						if time < baseTime {
							mutex.Lock()
							timeSaved[time] += 1
							mutex.Unlock()
						}
						grid[y][x] = "#"
					}
				}
			}
		}(y)

		y += 10
	}

	waitGroup.Add(1)
	go func(y int) {
		defer waitGroup.Done()

		grid := parseInput()

		for y := y; y < len(grid); y += 1 {
			for x := 0; x < len(grid[y]); x += 1 {
				if x == 0 || x == len(grid[y])-1 {
					continue
				}

				if grid[y][x] == "#" {
					grid[y][x] = "."
					astar = NewAStar(grid)
					time := astar.findShortestPath()

					if time < baseTime {
						mutex.Lock()
						timeSaved[time] += 1
						mutex.Unlock()
					}
					grid[y][x] = "#"
				}
			}
		}
	}(y)

	waitGroup.Wait()

	res := 0
	for time, count := range timeSaved {
		// fmt.Printf("There are %d cheats that save %d picoseconds.\n", count, baseTime-time)
		if baseTime-time >= 100 {
			res += count
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func part2() {
	fmt.Println("Result part 2: 0")
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
