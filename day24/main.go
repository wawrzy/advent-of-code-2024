package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Wire struct {
	Name      string
	Value     int
	Evaluated bool
}

type Gate struct {
	NameInput1 string
	NameInput2 string

	Input1 *Wire
	Input2 *Wire

	Output *Wire

	Operation string
}

func (g *Gate) Evaluate() {
	switch g.Operation {
	case "AND":
		g.Output.Value = g.Input1.Value & g.Input2.Value
		g.Output.Evaluated = true
	case "OR":
		g.Output.Value = g.Input1.Value | g.Input2.Value
		g.Output.Evaluated = true
	case "XOR":
		g.Output.Value = g.Input1.Value ^ g.Input2.Value
		g.Output.Evaluated = true
	default:
		panic("Unknown operation")
	}
}

type Wires map[string]*Wire

func parseInput() (Wires, []Gate) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(input), "\n\n")
	inputWires := strings.Split(parts[0], "\n")

	wires := make(Wires, 0)

	for _, part := range inputWires {
		p := strings.Split(part, ": ")

		value, err := strconv.Atoi(p[1])
		if err != nil {
			panic(err)
		}

		wires[p[0]] = &Wire{
			Name:      p[0],
			Value:     value,
			Evaluated: true,
		}
	}

	inputConnections := strings.Split(parts[1], "\n")

	gates := make([]Gate, 0)

	for _, part := range inputConnections {
		p := strings.Split(part, " ")

		output := &Wire{
			Name: p[4],
		}

		wires[p[4]] = output

		gate := Gate{
			NameInput1: p[0],
			NameInput2: p[2],

			Input1: wires[p[0]],
			Input2: wires[p[2]],

			Output: output,

			Operation: p[1],
		}

		gates = append(gates, gate)
	}
	return wires, gates
}

func part1() {
	wires, gates := parseInput()

	idx := 0
	nb := 0

	for {
		if idx >= len(gates) {
			if nb >= len(gates) {
				break
			} else {
				idx = 0
				nb = 0
			}
		}

		gate := gates[idx]

		if gate.Input1 == nil || gate.Input2 == nil {
			idx++
			continue
		}

		if gate.Input1.Evaluated == false || gate.Input2.Evaluated == false {
			idx++
			continue
		}

		gate.Evaluate()

		for i, g := range gates {
			if g.NameInput1 == gate.Output.Name {
				gates[i].Input1 = gate.Output
			}

			if g.NameInput2 == gate.Output.Name {
				gates[i].Input2 = gate.Output
			}
		}

		nb++
		idx++
	}

	res := 0

	for _, wire := range wires {
		if wire.Name[0] == 'z' {
			bitPosition, err := strconv.Atoi(wire.Name[1:])
			if err != nil {
				panic(err)
			}

			res |= wire.Value << bitPosition
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func getLastGate(wires Wires, gates []Gate) Gate {
	l := make([]string, 0)

	for _, wire := range wires {
		if strings.HasPrefix(wire.Name, "z") {
			l = append(l, wire.Name)
		}
	}

	slices.Sort(l)

	for _, gate := range gates {
		if gate.Output.Name == l[len(l)-1] {
			return gate
		}
	}

	panic("No gate found")
}

func isInputWire(name string) bool {
	return name[0] == 'x' || name[0] == 'y'
}

func areInputsFirstBit(name1, name2 string) bool {
	return name1[1:] == "00" && name2[1:] == "00"
}

func part2() {
	wires, gates := parseInput()

	lastGate := getLastGate(wires, gates)

	res := make([]string, 0)

	for _, gate := range gates {
		isFaulty := false

		if strings.HasPrefix(gate.Output.Name, "z") && gate != lastGate {
			isFaulty = gate.Operation != "XOR"
		} else if !strings.HasPrefix(gate.Output.Name, "z") && !isInputWire(gate.NameInput1) && !isInputWire(gate.NameInput2) {
			isFaulty = gate.Operation == "XOR"
		} else if isInputWire(gate.NameInput1) && isInputWire(gate.NameInput2) && !areInputsFirstBit(gate.NameInput1, gate.NameInput2) {
			nextShouldBeOr := gate.Operation == "AND"
			nextShouldBeXor := gate.Operation == "XOR"
			nextShouldBeAnd := gate.Operation == "XOR"

			for _, other := range gates {
				if other != gate && (other.NameInput1 == gate.Output.Name || other.NameInput2 == gate.Output.Name) {
					if nextShouldBeOr && other.Operation == "OR" {
						nextShouldBeOr = false
					}
					if nextShouldBeXor && other.Operation == "XOR" {
						nextShouldBeXor = false
					}
					if nextShouldBeAnd && other.Operation == "AND" {
						nextShouldBeAnd = false
					}
				}
			}

			isFaulty = nextShouldBeOr || nextShouldBeXor || nextShouldBeAnd
		}

		if isFaulty {
			res = append(res, gate.Output.Name)
		}
	}

	slices.Sort(res)

	fmt.Printf("Result part 2: %s\n", strings.Join(res, ","))
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
