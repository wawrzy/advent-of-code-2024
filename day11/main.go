package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func parseInput() []int {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	values := make([]int, 0)

	for _, n := range strings.Split(string(input), " ") {
		value, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		values = append(values, value)
	}

	return values
}

var (
	memo      = make(map[[2]int]int)
	memoMutex sync.Mutex
)

func blink(value int, nb int) int {
	memoMutex.Lock()
	if result, found := memo[[2]int{value, nb}]; found {
		memoMutex.Unlock()
		return result
	}
	memoMutex.Unlock()

	if nb == 0 {
		return 0
	}

	if value == 0 {
		return blink(1, nb-1)
	}

	isEven := true
	l := 0

	v := value
	for v > 0 {
		v /= 10
		isEven = !isEven
		l++
	}

	var result int
	if isEven {
		half := 1
		for i := 0; i < l/2; i++ {
			half *= 10
		}
		firstHalf := value / half
		secondHalf := value % half

		result = blink(firstHalf, nb-1) + blink(secondHalf, nb-1) + 1
	} else {
		result = blink(value*2024, nb-1)
	}

	memoMutex.Lock()
	memo[[2]int{value, nb}] = result
	memoMutex.Unlock()

	return result
}

func blinks(maxBlink int) int {
	values := parseInput()

	res := 0

	var waitGroup sync.WaitGroup

	for _, value := range values {
		waitGroup.Add(1)
		go func(value int) {
			defer waitGroup.Done()
			res += blink(value, maxBlink) + 1
		}(value)
	}

	waitGroup.Wait()

	return res
}

func part1() {
	fmt.Printf("Result part 1: %d\n", blinks(25))
}

func part2() {
	fmt.Printf("Result part 2: %d\n", blinks(75))
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
