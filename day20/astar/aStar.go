package astar

import (
	"math"
	"slices"
)

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

func (a *AStar) FindShortestPath() (int, [][2]int) {
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

	// Backtrack
	path := make([][2]int, 0)
	endIdx := len(a.closeList) - 1
	for i := endIdx; i >= 0; i-- {
		// if a.closeList[i][0] == a.StartX && a.closeList[i][1] == a.StartY {
		// 	break
		// }
		path = append(path, a.closeList[i])
	}

	return a.Grid[a.EndY][a.EndX].F, path
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
