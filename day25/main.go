package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type (
	Lock []int
	Key  []int
)

func (l Lock) IsKeyFit(k Key) bool {
	for i := 0; i < 5; i++ {
		if l[i]+k[i] > 5 {
			return false
		}
	}

	return true
}

func isLock(entity []string) bool {
	return entity[0] == "#####" && entity[len(entity)-1] == "....."
}

func countPins(entity []string) []int {
	pins := make([]int, 0)

	for x := 0; x < 5; x++ {
		count := 0

		for y := 1; y < len(entity)-1; y++ {
			if entity[y][x] == '#' {
				count++
			}
		}

		pins = append(pins, count)
	}

	return pins
}

func parseInput() ([]Lock, []Key) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	locks := make([]Lock, 0)
	keys := make([]Key, 0)

	parts := strings.Split(string(input), "\n\n")

	for _, part := range parts {
		entity := strings.Split(part, "\n")

		if isLock(entity) {
			locks = append(locks, countPins(entity))
		} else {
			keys = append(keys, countPins(entity))
		}
	}

	return locks, keys
}

func solve() {
	locks, keys := parseInput()
	res := 0

	for _, lock := range locks {
		for _, key := range keys {
			if lock.IsKeyFit(key) {
				res++
			}
		}
	}

	fmt.Printf("Result: %d\n", res)
}

func main() {
	now := time.Now()
	solve()
	duration := time.Since(now)

	fmt.Printf("Duration: %v\n", duration)
}
