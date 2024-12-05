package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Rules map[int][]int

func parseInput() (Rules, [][]int) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(input), "\n\n")

	if len(parts) != 2 {
		panic("Invalid input")
	}

	part1 := strings.Split(parts[0], "\n")
	part2 := strings.Split(parts[1], "\n")

	part1Parsed := make(Rules, len(part1))
	for _, line := range part1 {
		splited := strings.Split(line, "|")

		n1, err := strconv.Atoi(splited[0])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(splited[1])
		if err != nil {
			panic(err)
		}

		_, ok := part1Parsed[n1]
		if ok {
			part1Parsed[n1] = append(part1Parsed[n1], n2)
		} else {
			part1Parsed[n1] = []int{n2}
		}
	}

	part2Parsed := make([][]int, 0, len(part2))
	for _, line := range part2 {
		splited := strings.Split(line, ",")

		updates := make([]int, 0, len(splited))
		for _, n := range splited {
			n, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}

			updates = append(updates, n)
		}

		part2Parsed = append(part2Parsed, updates)
	}

	return part1Parsed, part2Parsed
}

func isOrdered(rules Rules, update []int) (bool, int) {
	current := update[0]
	errorIndex := -1

	for i := 0; i < len(update)-1; i++ {
		next := update[i+1]

		_, ok := rules[current]
		if !ok {
			errorIndex = i
			break
		}

		if slices.Index(rules[current], next) == -1 {
			errorIndex = i
			break
		}

		current = next
	}

	return errorIndex == -1, errorIndex
}

func part1() {
	rules, updates := parseInput()

	res := 0

	for _, update := range updates {
		ordered, _ := isOrdered(rules, update)

		if ordered {
			res += update[len(update)/2]
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func part2() {
	rules, updates := parseInput()

	res := 0

	for _, update := range updates {
		ordered, _ := isOrdered(rules, update)

		if !ordered {
			for {
				ok, errorIndex := isOrdered(rules, update)
				if ok {
					break
				}

				// move the error index to the end
				v := update[errorIndex]
				update = append(update[:errorIndex], update[errorIndex+1:]...)
				update = append(update, v)
			}

			res += update[len(update)/2]
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
