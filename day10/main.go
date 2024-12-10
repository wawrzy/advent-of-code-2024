package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Map [][]int

func parseInput() Map {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	topographicMap := make(Map, 0)

	for _, line := range strings.Split(string(input), "\n") {
		lineMap := make([]int, 0)

		for _, char := range line {
			value, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}

			lineMap = append(lineMap, value)
		}

		topographicMap = append(topographicMap, lineMap)
	}

	return topographicMap
}

type Positions map[[2]int]int

func countHikingTrails(topographicMap Map, posX, posY int, positions *Positions) {
	currentValue := topographicMap[posX][posY]

	if currentValue == 9 {
		(*positions)[[2]int{posX, posY}] += 1
		return
	}

	if posX > 0 && topographicMap[posX-1][posY] == currentValue+1 {
		countHikingTrails(topographicMap, posX-1, posY, positions)
	}
	if posY > 0 && topographicMap[posX][posY-1] == currentValue+1 {
		countHikingTrails(topographicMap, posX, posY-1, positions)
	}
	if posX < len(topographicMap)-1 && topographicMap[posX+1][posY] == currentValue+1 {
		countHikingTrails(topographicMap, posX+1, posY, positions)
	}
	if posY < len(topographicMap[0])-1 && topographicMap[posX][posY+1] == currentValue+1 {
		countHikingTrails(topographicMap, posX, posY+1, positions)
	}
}

func part1() {
	topographicMap := parseInput()

	res := 0

	for x, line := range topographicMap {
		for y, value := range line {
			if value == 0 {
				positions := make(Positions, 0)
				countHikingTrails(topographicMap, x, y, &positions)
				res += len(positions)
			}
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func part2() {
	topographicMap := parseInput()

	res := 0

	for x, line := range topographicMap {
		for y, value := range line {
			if value == 0 {
				positions := make(Positions, 0)
				countHikingTrails(topographicMap, x, y, &positions)

				for _, v := range positions {
					res += v
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
