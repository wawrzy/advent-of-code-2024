package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseInput() []int {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	secrets := make([]int, 0)

	for _, line := range lines {
		secret, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		secrets = append(secrets, secret)
	}

	return secrets
}

type Secret int

func (s Secret) mix(a, b int) int {
	return a ^ b
}

func (s Secret) prune(value int) int {
	return value % 16777216
}

func (s Secret) Generate(depth int) (int, [][2]int) {
	value := int(s)

	diffs := make([][2]int, 0)
	prev := value % 10

	diffs = append(diffs, [2]int{value % 10, 0})

	for i := 0; i < depth; i++ {
		value = s.mix(value*64, value)
		value = s.prune(value)

		value = s.mix(int(math.Floor(float64(value)/32)), value)
		value = s.prune(value)

		value = s.mix(value*2048, value)
		value = s.prune(value)

		diffs = append(diffs, [2]int{value%10 - prev, value % 10})
		prev = value % 10
	}

	return int(value), diffs
}

func part1() {
	secrets := parseInput()

	res := 0

	for _, s := range secrets {
		secret := Secret(s)
		r, _ := secret.Generate(2000)

		res += r
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func part2() {
	secrets := parseInput()

	sequences := make(map[[4]int]map[int]int)

	for buyerIdx, s := range secrets {
		secret := Secret(s)
		_, diffs := secret.Generate(2000)

		for i := 0; i < len(diffs)-3; i++ {
			seq := [4]int{diffs[i][0], diffs[i+1][0], diffs[i+2][0], diffs[i+3][0]}

			if _, ok := sequences[seq]; !ok {
				sequences[seq] = make(map[int]int)
			}

			if _, ok := sequences[seq][buyerIdx]; ok {
				continue
			}

			sequences[seq][buyerIdx] = diffs[i+3][1]
		}
	}

	max := 0

	for _, s := range sequences {
		total := 0
		for _, v := range s {
			total += v
		}

		if total > max {
			max = total
		}
	}

	fmt.Printf("Result part 2: %d\n", max)
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
