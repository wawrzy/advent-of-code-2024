package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseInput() []string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	codes := strings.Split(string(input), "\n")

	return codes
}

const (
	digitKeys = "123456789A"
)

var pathsR = map[[2]rune]string{
	{'A', '0'}: "<A",
	{'0', 'A'}: ">A",
	{'A', '1'}: "^<<A",
	{'1', 'A'}: ">>vA",
	{'A', '2'}: "<^A",
	{'2', 'A'}: "v>A",
	{'A', '3'}: "^A",
	{'3', 'A'}: "vA",
	{'A', '4'}: "^^<<A",
	{'4', 'A'}: ">>vvA",
	{'A', '5'}: "<^^A",
	{'5', 'A'}: "vv>A",
	{'A', '6'}: "^^A",
	{'6', 'A'}: "vvA",
	{'A', '7'}: "^^^<<A",
	{'7', 'A'}: ">>vvvA",
	{'A', '8'}: "<^^^A",
	{'8', 'A'}: "vvv>A",
	{'A', '9'}: "^^^A",
	{'9', 'A'}: "vvvA",
	{'0', '1'}: "^<A",
	{'1', '0'}: ">vA",
	{'0', '2'}: "^A",
	{'2', '0'}: "vA",
	{'0', '3'}: "^>A",
	{'3', '0'}: "<vA",
	{'0', '4'}: "^<^A",
	{'4', '0'}: ">vvA",
	{'0', '5'}: "^^A",
	{'5', '0'}: "vvA",
	{'0', '6'}: "^^>A",
	{'6', '0'}: "<vvA",
	{'0', '7'}: "^^^<A",
	{'7', '0'}: ">vvvA",
	{'0', '8'}: "^^^A",
	{'8', '0'}: "vvvA",
	{'0', '9'}: "^^^>A",
	{'9', '0'}: "<vvvA",
	{'1', '2'}: ">A",
	{'2', '1'}: "<A",
	{'1', '3'}: ">>A",
	{'3', '1'}: "<<A",
	{'1', '4'}: "^A",
	{'4', '1'}: "vA",
	{'1', '5'}: "^>A",
	{'5', '1'}: "<vA",
	{'1', '6'}: "^>>A",
	{'6', '1'}: "<<vA",
	{'1', '7'}: "^^A",
	{'7', '1'}: "vvA",
	{'1', '8'}: "^^>A",
	{'8', '1'}: "<vvA",
	{'1', '9'}: "^^>>A",
	{'9', '1'}: "<<vvA",
	{'2', '3'}: ">A",
	{'3', '2'}: "<A",
	{'2', '4'}: "<^A",
	{'4', '2'}: "v>A",
	{'2', '5'}: "^A",
	{'5', '2'}: "vA",
	{'2', '6'}: "^>A",
	{'6', '2'}: "<vA",
	{'2', '7'}: "<^^A",
	{'7', '2'}: "vv>A",
	{'2', '8'}: "^^A",
	{'8', '2'}: "vvA",
	{'2', '9'}: "^^>A",
	{'9', '2'}: "<vvA",
	{'3', '4'}: "<<^A",
	{'4', '3'}: "v>>A",
	{'3', '5'}: "<^A",
	{'5', '3'}: "v>A",
	{'3', '6'}: "^A",
	{'6', '3'}: "vA",
	{'3', '7'}: "<<^^A",
	{'7', '3'}: "vv>>A",
	{'3', '8'}: "<^^A",
	{'8', '3'}: "vv>A",
	{'3', '9'}: "^^A",
	{'9', '3'}: "vvA",
	{'4', '5'}: ">A",
	{'5', '4'}: "<A",
	{'4', '6'}: ">>A",
	{'6', '4'}: "<<A",
	{'4', '7'}: "^A",
	{'7', '4'}: "vA",
	{'4', '8'}: "^>A",
	{'8', '4'}: "<vA",
	{'4', '9'}: "^>>A",
	{'9', '4'}: "<<vA",
	{'5', '6'}: ">A",
	{'6', '5'}: "<A",
	{'5', '7'}: "<^A",
	{'7', '5'}: "v>A",
	{'5', '8'}: "^A",
	{'8', '5'}: "vA",
	{'5', '9'}: "^>A",
	{'9', '5'}: "<vA",
	{'6', '7'}: "<<^A",
	{'7', '6'}: "v>>A",
	{'6', '8'}: "<^A",
	{'8', '6'}: "v>A",
	{'6', '9'}: "^A",
	{'9', '6'}: "vA",
	{'7', '8'}: ">A",
	{'8', '7'}: "<A",
	{'7', '9'}: ">>A",
	{'9', '7'}: "<<A",
	{'8', '9'}: ">A",
	{'9', '8'}: "<A",
	{'<', '^'}: ">^A",
	{'^', '<'}: "v<A",
	{'<', 'v'}: ">A",
	{'v', '<'}: "<A",
	{'<', '>'}: ">>A",
	{'>', '<'}: "<<A",
	{'<', 'A'}: ">>^A",
	{'A', '<'}: "v<<A",
	{'^', 'v'}: "vA",
	{'v', '^'}: "^A",
	{'^', '>'}: "v>A",
	{'>', '^'}: "<^A",
	{'^', 'A'}: ">A",
	{'A', '^'}: "<A",
	{'v', '>'}: ">A",
	{'>', 'v'}: "<A",
	{'v', 'A'}: "^>A",
	{'A', 'v'}: "<vA",
	{'>', 'A'}: "^A",
	{'A', '>'}: "vA",
}

var keypadDigits = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{" ", "0", "A"},
}

var keypadDirections = [][]string{
	{" ", "^", "A"},
	{"<", "v", ">"},
}

var cache = make(map[string]int)

func solve(code string, depth int) int {
	if depth < 0 {
		return len(code) - 1
	}

	cacheKey := code + "-" + strconv.Itoa(depth)

	if val, ok := cache[cacheKey]; ok {
		return val
	}

	total := 0

	for i := 1; i < len(code); i++ {
		prev := rune(code[i-1])
		curr := rune(code[i])

		p := [2]rune{prev, curr}

		if prev == curr {
			total += 1
		} else {
			total += solve("A"+pathsR[p], depth-1)
		}
	}

	cache[cacheKey] = total

	return total
}

func part1(depth int) {
	codes := parseInput()
	res := 0

	for _, code := range codes {
		cache = make(map[string]int)
		codeVal, err := strconv.Atoi(code[:3])
		if err != nil {
			panic(err)
		}

		r := solve("A"+code, depth)

		res += codeVal * r
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func part2(depth int) {
	codes := parseInput()
	res := 0

	for _, code := range codes {
		cache = make(map[string]int)
		codeVal, err := strconv.Atoi(code[:3])
		if err != nil {
			panic(err)
		}

		r := solve("A"+code, depth)

		res += codeVal * r
	}

	fmt.Printf("Result part 2: %d\n", res)
}

func main() {
	now := time.Now()
	part1(2)
	duration := time.Since(now)

	fmt.Printf("Duration part 1: %v\n", duration)

	now = time.Now()
	part2(25)
	duration = time.Since(now)

	fmt.Printf("Duration part 2: %v\n", duration)
}
