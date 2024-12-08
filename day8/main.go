package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func parseInput() [][]string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid := make([][]string, 0)

	for _, line := range strings.Split(string(input), "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func part1() {
	grid := parseInput()

	antennas := make(map[string][][2]int)
	coords := make(map[[2]int]string)

	for x, row := range grid {
		for y, cell := range row {
			if cell != "." {
				coords[[2]int{x, y}] = cell
				antennas[cell] = append(antennas[cell], [2]int{x, y})
			}
		}
	}

	res := 0

	for _, locations := range antennas {
		for _, locA := range locations {
			for _, locB := range locations {
				if locA != locB {
					antinodeX := locA[0] + (locA[0] - locB[0])
					antinodeY := locA[1] + (locA[1] - locB[1])

					if antinodeX >= 0 && antinodeX < len(grid) && antinodeY >= 0 && antinodeY < len(grid[0]) {
						if grid[antinodeX][antinodeY] != "#" {
							grid[antinodeX][antinodeY] = "#"
							res++
						}
					}
				}
			}
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func part2() {
	grid := parseInput()

	antennas := make(map[string][][2]int)

	for x, row := range grid {
		for y, cell := range row {
			if cell != "." {
				antennas[cell] = append(antennas[cell], [2]int{x, y})
			}
		}
	}

	res := 0

	for _, locations := range antennas {
		for _, locA := range locations {
			for _, locB := range locations {
				if locA != locB {
					antinodeX := locA[0] + (locA[0] - locB[0])
					antinodeY := locA[1] + (locA[1] - locB[1])

					for {
						if antinodeX >= 0 && antinodeX < len(grid) && antinodeY >= 0 && antinodeY < len(grid[0]) {
							if grid[antinodeX][antinodeY] != "#" {
								grid[antinodeX][antinodeY] = "#"
								res++
							}

							antinodeX += (locA[0] - locB[0])
							antinodeY += (locA[1] - locB[1])
						} else {
							break
						}
					}
				}
			}
		}
	}

	for _, row := range grid {
		for _, cell := range row {
			if cell != "." && cell != "#" && len(antennas[cell]) > 1 {
				res++
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
