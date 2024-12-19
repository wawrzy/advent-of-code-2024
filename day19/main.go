package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func parseInput() ([]string, []string) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(input), "\n\n")

	patterns := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1], "\n")

	return patterns, designs
}

func canBuildDesign(patterns []string, design string, current string, tried map[string]bool, depth int) bool {
	if design == "" {
		return true
	}

	if tried[current] {
		return false
	}

	// fmt.Println()
	// fmt.Printf("Design: %s\n", design)

	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) && len(design) >= len(pattern) {
			// fmt.Printf("Pattern: %s, Design: %s, Depth: %d\n", pattern, current+pattern, depth)
			// fmt.Printf("canBuildDesign(_, %s, %s)\n", design[len(pattern):], current+pattern)

			if canBuildDesign(patterns, design[len(pattern):], current+pattern, tried, depth+1) {
				return true
			}

			tried[current+pattern] = true
		}
	}

	return false
}

func getAllOptionsForDesign(patterns []string, design string, current string, tried map[string]bool, options map[string]int) int {
	if design == "" {
		return 1
	}

	if tried[current] {
		return 0
	}

	if resOptions, ok := options[current]; ok {
		return resOptions
	}

	allOptions := 0

	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) && len(design) >= len(pattern) {
			if nbOptions := getAllOptionsForDesign(patterns, design[len(pattern):], current+pattern, tried, options); nbOptions > 0 {
				if _, ok := options[current+pattern]; !ok {
					options[current+pattern] = nbOptions
				}
				allOptions += nbOptions
			} else {
				tried[current+pattern] = true
			}
		}
	}

	return allOptions
}

func part1() {
	patterns, designs := parseInput()

	res := 0

	for _, design := range designs {
		tried := make(map[string]bool)
		if canBuildDesign(patterns, design, "", tried, 0) {
			res++
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func part2() {
	patterns, designs := parseInput()

	res := 0

	for _, design := range designs {
		tried := make(map[string]bool)

		if !canBuildDesign(patterns, design, "", tried, 0) {
			continue
		}

		tried = make(map[string]bool)
		options := make(map[string]int)

		res += getAllOptionsForDesign(patterns, design, "", tried, options)
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
