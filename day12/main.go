package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Plot struct {
	PlantType string
	Region    int
	Perimeter int
	Counted   bool
	X         int
	Y         int
}

type Garden [][]Plot

func parseInput() Garden {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	garden := make(Garden, 0)

	for _, line := range strings.Split(string(input), "\n") {
		gardenLine := make([]Plot, 0, len(line))

		for _, c := range line {
			gardenLine = append(gardenLine, Plot{PlantType: string(c)})
		}

		garden = append(garden, gardenLine)
	}

	return garden
}

func countRegion(garden Garden, region []Plot, x, y int) (Garden, []Plot) {
	plant := garden[x][y].PlantType

	garden[x][y].X = x
	garden[x][y].Y = y
	garden[x][y].Counted = true

	perimeter := 0

	if x > 0 {
		if garden[x-1][y].PlantType != plant {
			perimeter += 1
		} else if !garden[x-1][y].Counted {
			garden, region = countRegion(garden, region, x-1, y)
		}
	} else {
		perimeter += 1
	}

	if x < len(garden)-1 {
		if garden[x+1][y].PlantType != plant {
			perimeter += 1
		} else if !garden[x+1][y].Counted {
			garden, region = countRegion(garden, region, x+1, y)
		}
	} else {
		perimeter += 1
	}

	if y > 0 {
		if garden[x][y-1].PlantType != plant {
			perimeter += 1
		} else if !garden[x][y-1].Counted {
			garden, region = countRegion(garden, region, x, y-1)
		}
	} else {
		perimeter += 1
	}

	if y < len(garden[x])-1 {
		if garden[x][y+1].PlantType != plant {
			perimeter += 1
		} else if !garden[x][y+1].Counted {
			garden, region = countRegion(garden, region, x, y+1)
		}
	} else {
		perimeter += 1
	}

	garden[x][y].Perimeter = perimeter

	region = append(region, garden[x][y])

	return garden, region
}

func part1() {
	garden := parseInput()

	res := 0

	for x := range garden {
		for y := range garden[x] {
			if garden[x][y].Counted {
				continue
			}

			region := make([]Plot, 0)
			garden, region = countRegion(garden, region, x, y)

			area := len(region)
			perimeter := 0
			for _, plot := range region {
				perimeter += plot.Perimeter
			}

			price := area * perimeter
			res += price
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func countRowSides(tmp [][]string) int {
	res := 0

	sides := make([][]int, 0, len(tmp))

	for x := 1; x < len(tmp)-1; x++ {
		side := make([]int, 0, len(tmp[x]))
		isInRegion := false

		for y := 0; y < len(tmp[x]); y++ {
			if tmp[x][y] != "." && !isInRegion {
				isInRegion = true
				side = append(side, y)
			} else if tmp[x][y] == "." && isInRegion {
				isInRegion = false
				side = append(side, y-1)
			} else {
				side = append(side, -1)
			}
		}

		sides = append(sides, side)
	}

	for {
		prev := -1
		change := false

		for i, side := range sides {
			if len(side) == 0 {
				// prev = -1
				continue
			}
			change = true

			if side[0] == -1 {
				sides[i] = side[1:]
				prev = -1
				continue
			}

			if side[0] != prev {
				res += 1
				prev = side[0]
			}
			sides[i] = side[1:]
		}

		if !change {
			break
		}
	}

	return res
}

func countColSides(tmp [][]string) int {
	res := 0

	sides := make([][]int, 0, len(tmp[0]))

	for y := 1; y < len(tmp[0])-1; y++ {
		side := make([]int, 0, len(tmp))
		isInRegion := false

		for x := 0; x < len(tmp); x++ {
			if tmp[x][y] != "." && !isInRegion {
				isInRegion = true
				side = append(side, x)
			} else if tmp[x][y] == "." && isInRegion {
				isInRegion = false
				side = append(side, x-1)
			} else {
				side = append(side, -1)
			}
		}

		sides = append(sides, side)
	}

	for {
		change := false
		prev := -1

		for i, side := range sides {
			if len(side) == 0 {
				continue
			}
			change = true

			if side[0] == -1 {
				sides[i] = side[1:]
				prev = -1
				continue
			}

			if side[0] != prev {
				res += 1
				prev = side[0]
			}
			sides[i] = side[1:]
		}

		if !change {
			break
		}
	}

	return res
}

func countSides(region []Plot) int {
	minX, minY := 1000000, 1000000
	maxX, maxY := 0, 0

	for _, plot := range region {
		if plot.X < minX {
			minX = plot.X
		}
		if plot.Y < minY {
			minY = plot.Y
		}
		if plot.X > maxX {
			maxX = plot.X
		}
		if plot.Y > maxY {
			maxY = plot.Y
		}
	}

	rows := maxX - minX + 3
	cols := maxY - minY + 3

	tmp := make([][]string, 0, rows)
	for i := 0; i < rows; i++ {
		tmp = append(tmp, make([]string, 0, cols))
		for j := 0; j < cols; j++ {
			tmp[i] = append(tmp[i], ".")
		}
	}
	for _, plot := range region {
		tmp[plot.X-minX+1][plot.Y-minY+1] = plot.PlantType
	}

	rowSides := countRowSides(tmp)
	colSides := countColSides(tmp)

	return rowSides + colSides
}

func part2() {
	garden := parseInput()

	res := 0

	for x := range garden {
		for y := range garden[x] {
			if garden[x][y].Counted {
				continue
			}
			region := make([]Plot, 0)
			garden, region = countRegion(garden, region, x, y)

			sides := countSides(region)
			area := len(region)

			res += area * sides
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
