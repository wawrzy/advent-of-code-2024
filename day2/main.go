package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseInput() [][]int {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	reports := make([][]int, 0, len(lines))

	for _, line := range lines {
		coords := strings.Split(line, " ")

		if len(coords) > 1 {
			report := make([]int, 0, len(coords))

			for _, coord := range coords {
				n, err := strconv.Atoi(coord)
				if err != nil {
					panic(err)
				}

				report = append(report, n)
			}

			reports = append(reports, report)
		}
	}

	return reports
}

func part1() {
	reports := parseInput()

	res := 0

	for _, report := range reports {
		isIncreasing := report[0] < report[1]
		isDecreasing := report[0] > report[1]

		if !isIncreasing && !isDecreasing {
			continue
		}

		isValid := true

		for i := 0; i < len(report)-1; i++ {
			diff := 0
			if isIncreasing {
				diff = report[i+1] - report[i]
			} else {
				diff = report[i] - report[i+1]
			}

			if diff < 1 || diff > 3 {
				isValid = false
				break
			}
		}

		if isValid {
			res++
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func isReportValid(report []int, allowBadLevel bool) bool {
	isValid := true
	prev := report[0]
	hasBadLevel := false

	isIncreasing := report[0] < report[1]
	isDecreasing := report[0] > report[1]

	if !isIncreasing && !isDecreasing {
		return false
	}

	for i := 1; i < len(report); i++ {
		diff := report[i] - prev

		hasError := false

		if isDecreasing {
			hasError = diff >= 0
			diff = int(math.Abs(float64(diff)))
		} else {
			hasError = diff <= 0
		}

		if diff >= 1 && diff <= 3 && !hasError {
			prev = report[i]
			continue
		} else if !hasBadLevel && allowBadLevel {
			hasBadLevel = true
		} else {
			isValid = false
			break
		}
	}

	return isValid
}

func part2() {
	reports := parseInput()

	res := 0

	for _, report := range reports {
		isValid := isReportValid(report, true)

		if !isValid {
			isValid = isReportValid(report[1:], false)
		}

		if isValid {
			res++
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
