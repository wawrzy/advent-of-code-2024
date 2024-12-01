package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func parseInput() ([]int, []int) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	pair1 := make([]int, 0, len(lines))
	pair2 := make([]int, 0, len(lines))

	for _, line := range lines {
		coords := strings.Split(line, "   ")

		if len(coords) == 2 {
			n1, err := strconv.Atoi(coords[0])
			if err != nil {
				panic(err)
			}
			n2, err := strconv.Atoi(coords[1])
			if err != nil {
				panic(err)
			}

			pair1 = append(pair1, n1)
			pair2 = append(pair2, n2)
		}
	}

	return pair1, pair2
}

func part1() {
	pair1, pair2 := parseInput()

	sort.Ints(pair1)
	sort.Ints(pair2)

	res := 0

	for i := 0; i < len(pair1); i++ {
		res += int(math.Abs(float64(pair2[i]) - float64(pair1[i])))
	}

	fmt.Printf("Result part 1: %d\n", res)
}

type SimilarityScore struct {
	Exist bool
	Count int
	Occur int
}

func part2() {
	pair1, pair2 := parseInput()

	scores := make(map[int]SimilarityScore)

	for i := 0; i < len(pair1); i++ {
		score1, ok := scores[pair1[i]]

		if ok {
			score1.Exist = true
			score1.Occur++
			scores[pair1[i]] = score1
		} else {
			scores[pair1[i]] = SimilarityScore{Exist: true, Occur: 1}
		}

		score2, ok := scores[pair2[i]]
		if ok {
			score2.Count++
			scores[pair2[i]] = score2
		} else {
			scores[pair2[i]] = SimilarityScore{Count: 1}
		}
	}

	res := 0

	for key, score := range scores {
		if score.Exist {
			res += key * score.Occur * score.Count
		}
	}

	fmt.Printf("Result part 2: %v\n", res)
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
