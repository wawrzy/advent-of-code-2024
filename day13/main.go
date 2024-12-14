package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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

type ClawMachine struct {
	ButtonA [2]int
	ButtonB [2]int
	Prize   [2]int
}

func parseInput() []ClawMachine {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	machines := make([]ClawMachine, 0)

	for _, m := range strings.Split(string(input), "\n\n") {
		parts := strings.Split(m, "\n")

		values := strings.Split(parts[0], " ")[2:]
		buttonAx := strings.Split(values[0][:len(values[0])-1], "+")[1]
		buttonAy := strings.Split(values[1], "+")[1]

		values = strings.Split(parts[1], " ")[2:]
		buttonBx := strings.Split(values[0][:len(values[0])-1], "+")[1]
		buttonBy := strings.Split(values[1], "+")[1]

		values = strings.Split(parts[2], " ")[1:]
		prizeX := strings.Split(values[0][:len(values[0])-1], "=")[1]
		prizeY := strings.Split(values[1], "=")[1]

		aX, _ := strconv.Atoi(buttonAx)
		aY, _ := strconv.Atoi(buttonAy)
		bX, _ := strconv.Atoi(buttonBx)
		bY, _ := strconv.Atoi(buttonBy)
		pX, _ := strconv.Atoi(prizeX)
		pY, _ := strconv.Atoi(prizeY)

		machines = append(machines, ClawMachine{
			ButtonA: [2]int{aX, aY},
			ButtonB: [2]int{bX, bY},
			Prize:   [2]int{pX, pY},
		})
	}

	return machines
}

func resolve(xA, yA, xB, yB, prizeA, prizeB int) (int, int) {
	a := yA * xB
	b := xA * yB
	p := int(math.Abs(float64(prizeA*xB - prizeB*xA)))

	y := int(math.Abs(float64(a - b)))
	y = p / y

	x := (prizeA - yA*y) / xA

	return x, y
}

func part1() {
	machines := parseInput()

	res := 0

	for _, m := range machines {
		x, y := resolve(m.ButtonA[0], m.ButtonB[0], m.ButtonA[1], m.ButtonB[1], m.Prize[0], m.Prize[1])

		if x*m.ButtonA[0]+y*m.ButtonB[0] == m.Prize[0] && x*m.ButtonA[1]+y*m.ButtonB[1] == m.Prize[1] {
			res += x*3 + y
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func part2() {
	machines := parseInput()

	res := 0

	for _, m := range machines {
		prizeA := m.Prize[0] + 10000000000000
		prizeB := m.Prize[1] + 10000000000000

		x, y := resolve(m.ButtonA[0], m.ButtonB[0], m.ButtonA[1], m.ButtonB[1], prizeA, prizeB)

		if x*m.ButtonA[0]+y*m.ButtonB[0] == prizeA && x*m.ButtonA[1]+y*m.ButtonB[1] == prizeB {
			res += x*3 + y
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
