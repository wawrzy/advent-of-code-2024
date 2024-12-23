package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func parseInput() [][]string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	connections := make([][]string, 0)

	for _, line := range lines {
		computers := strings.Split(line, "-")
		connections = append(connections, computers)
	}

	return connections
}

type Computer struct {
	Name string

	ConnectedTo map[string]*Computer
}

func part1() {
	connections := parseInput()

	computers := make(map[string]*Computer)

	interconnectedComputers := make([][]string, 0)

	for _, connection := range connections {
		computer1, ok := computers[connection[0]]
		if !ok {
			computer1 = &Computer{Name: connection[0]}
			computer1.ConnectedTo = make(map[string]*Computer)
			computers[connection[0]] = computer1
		}

		computer2, ok := computers[connection[1]]
		if !ok {
			computer2 = &Computer{Name: connection[1]}
			computer2.ConnectedTo = make(map[string]*Computer)
			computers[connection[1]] = computer2
		}

		computer1.ConnectedTo[connection[1]] = computer2
		computer2.ConnectedTo[connection[0]] = computer1
	}

	for computerNameA, computer := range computers {
		for computerNameB, connectedComputerB := range computer.ConnectedTo {
			for computerNameC, connectedComputerC := range connectedComputerB.ConnectedTo {
				_, ok := connectedComputerC.ConnectedTo[computerNameA]

				if ok {
					r := []string{computerNameA, computerNameB, computerNameC}

					slices.Sort(r)

					found := false
					for _, interconnected := range interconnectedComputers {
						if interconnected[0] == r[0] && interconnected[1] == r[1] && interconnected[2] == r[2] {
							found = true
							break
						}
					}

					if !found {
						interconnectedComputers = append(interconnectedComputers, r)
					}
				}
			}
		}
	}

	res := 0

	for _, interconnected := range interconnectedComputers {
		for _, computerName := range interconnected {
			if computerName[0] == 't' {
				res++
				break
			}
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func getInterconnected(computer *Computer, current []string) []string {
	var res []string = current

	for _, c := range computer.ConnectedTo {

		found := true
		for _, name := range current {
			if _, ok := c.ConnectedTo[name]; !ok {
				found = false
				break
			}
		}

		if found {
			nextComputers := make([]string, 0)
			for _, name := range current {
				nextComputers = append(nextComputers, name)
			}

			nextComputers = append(nextComputers, c.Name)
			nextComputers = getInterconnected(c, nextComputers)

			if len(nextComputers) > len(res) {
				res = nextComputers
			} else if len(res) == 0 {
				res = nextComputers
			}
		} else {
			return res
		}
	}

	return res
}

func part2() {
	connections := parseInput()
	computers := make(map[string]*Computer)

	for _, connection := range connections {
		computer1, ok := computers[connection[0]]
		if !ok {
			computer1 = &Computer{Name: connection[0]}
			computer1.ConnectedTo = make(map[string]*Computer)
			computers[connection[0]] = computer1
		}

		computer2, ok := computers[connection[1]]
		if !ok {
			computer2 = &Computer{Name: connection[1]}
			computer2.ConnectedTo = make(map[string]*Computer)
			computers[connection[1]] = computer2
		}

		computer1.ConnectedTo[connection[1]] = computer2
		computer2.ConnectedTo[connection[0]] = computer1
	}

	max := 0
	var maxInterconnected []string

	for name := range computers {
		computer := computers[name]

		interconnected := getInterconnected(computer, []string{name})

		if len(interconnected) > max {
			max = len(interconnected)
			maxInterconnected = interconnected
		}
	}

	slices.Sort(maxInterconnected)

	fmt.Printf("Result part 2: %s\n", strings.Join(maxInterconnected, ","))
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
