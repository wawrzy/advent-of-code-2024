package main

import (
	"fmt"
	"math"
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

	graph := make([][]string, 0)

	parts := strings.Split(string(input), "\n")
	for _, part := range parts {
		row := make([]string, 0)
		for _, p := range strings.Split(part, "") {
			row = append(row, p)
		}
		graph = append(graph, row)
	}

	return graph
}

func popSmallest(queue [][2]int, dist map[[2]int]int) ([][2]int, [2]int) {
	min := math.MaxInt64
	minIndex := 0

	for i, q := range queue {
		if dist[q] < min {
			min = dist[q]
			minIndex = i
		}
	}

	res := queue[minIndex]

	if minIndex > 0 {
		queue = append(queue[:minIndex], queue[minIndex+1:]...)
	} else {
		queue = queue[1:]
	}

	return queue, res
}

func part1() {
	graph := parseInput()

	dist := make(map[[2]int]int)
	prev := make(map[[2]int][2]int)

	queue := make([][2]int, 0)

	for y, row := range graph {
		for x, cell := range row {
			if cell == "#" {
				continue
			}

			dist[[2]int{x, y}] = math.MaxInt64
			queue = append(queue, [2]int{x, y})

			if cell == "S" {
				dist[[2]int{x, y}] = 0
				prev[[2]int{x, y}] = [2]int{-1, -1}
			}
		}
	}

	var u [2]int
	end := [2]int{0, 0}

	for len(queue) > 0 {
		queue, u = popSmallest(queue, dist)

		if graph[u[1]][u[0]] == "E" {
			end = u
			break
		}

		directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

		dir := ""
		for _, d := range directions {
			if prev[u] == [2]int{u[0] + d[0], u[1] + d[1]} {
				if d[0] == 1 {
					dir = "L"
				}
				if d[0] == -1 {
					dir = "R"
				}
				if d[1] == 1 {
					dir = "U"
				}
				if d[1] == -1 {
					dir = "D"
				}
				break
			}
		}

		if dir == "" {
			dir = "R"
		}

		for _, d := range directions {
			v := [2]int{u[0] + d[0], u[1] + d[1]}

			if v[0] < 0 || v[0] >= len(graph[0]) || v[1] < 0 || v[1] >= len(graph) {
				continue
			}

			if graph[v[1]][v[0]] == "#" {
				continue
			}

			if !slices.Contains(queue, v) {
				continue
			}

			isTurn := false
			if dir == "L" && d[0] != -1 {
				isTurn = true
			} else if dir == "R" && d[0] != 1 {
				isTurn = true
			} else if dir == "U" && d[1] != -1 {
				isTurn = true
			} else if dir == "D" && d[1] != 1 {
				isTurn = true
			}

			alt := 0
			if isTurn {
				alt = dist[u] + 1001
			} else {
				alt = dist[u] + 1
			}

			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}

	fmt.Printf("Result part 1: %d\n", dist[end])
}

func main() {
	now := time.Now()
	part1()
	duration := time.Since(now)

	fmt.Printf("Duration part 1: %v\n", duration)
}
