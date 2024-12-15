package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type PositionType int

const (
	Wall PositionType = iota
	Block
	BlockLeft
	BlockRight
	Robot
	Empty
)

type Position struct {
	X    int
	Y    int
	Type PositionType
}

func (p Position) IsEmpty() bool {
	return p.Type == Empty
}

func (p Position) IsWall() bool {
	return p.Type == Wall
}

func (p Position) IsBlock() bool {
	return p.Type == Block || p.Type == BlockLeft || p.Type == BlockRight
}

func (p Position) IsBlockLeft() bool {
	return p.Type == BlockLeft
}

func (p Position) IsBlockRight() bool {
	return p.Type == BlockRight
}

func (p Position) IsRobot() bool {
	return p.Type == Robot
}

func (p Position) IsEqual(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

type WareHouse [][]Position

func (w WareHouse) GetRobotPosition() Position {
	for _, row := range w {
		for _, p := range row {
			if p.IsRobot() {
				return p
			}
		}
	}
	return Position{}
}

func (w WareHouse) GetScore() int {
	score := 0

	for _, row := range w {
		for _, p := range row {
			if p.IsBlockLeft() {
				currentScore := p.X + p.Y*100
				score += currentScore
			} else if p.IsBlockRight() {
				continue
			} else if p.IsBlock() {
				currentScore := p.X + p.Y*100
				score += currentScore
			}
		}
	}

	return score
}

func (w WareHouse) Display() {
	for _, row := range w {
		for _, p := range row {
			switch p.Type {
			case Wall:
				fmt.Printf("#")
			case Block:
				fmt.Printf("O")
			case BlockLeft:
				fmt.Printf("[")
			case BlockRight:
				fmt.Printf("]")
			case Robot:
				fmt.Printf("@")
			case Empty:
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func (w WareHouse) CanPushBlockRec(block Position, prev Position, addX int, addY int) bool {
	var left Position
	var right Position

	if block.IsBlockLeft() {
		left = block
		right = w[block.Y][block.X+1]
	} else {
		left = w[block.Y][block.X-1]
		right = block
	}

	nextPositionX := block.X + addX
	nextPositionY := block.Y + addY

	if nextPositionX < 0 || nextPositionX >= len(w[0]) || nextPositionY < 0 || nextPositionY >= len(w) {
		return false
	}

	nextPosition := w[nextPositionY][nextPositionX]

	if nextPosition.IsWall() {
		return false
	} else if nextPosition.IsEqual(left) || nextPosition.IsEqual(right) {
		return w.CanPushBlockRec(nextPosition, block, addX, addY)
	}

	if nextPosition.IsBlock() {
		ok := w.CanPushBlockRec(nextPosition, block, addX, addY)

		if !ok {
			return false
		}

		if block.IsBlockLeft() {
			if !prev.IsEqual(right) {
				return w.CanPushBlockRec(right, block, addX, addY)
			}
		} else {
			if !prev.IsEqual(left) {
				return w.CanPushBlockRec(left, block, addX, addY)
			}
		}

		return true
	}

	if nextPosition.IsEmpty() {
		if block.IsBlockLeft() {
			if !prev.IsEqual(right) {
				return w.CanPushBlockRec(right, block, addX, addY)
			}
		} else {
			if !prev.IsEqual(left) {
				return w.CanPushBlockRec(left, block, addX, addY)
			}
		}

		return true
	}

	return false
}

func (w WareHouse) PushBlockRec(block Position, prev Position, addX int, addY int) WareHouse {
	var left Position
	var right Position

	if block.IsBlockLeft() {
		left = block
		right = w[block.Y][block.X+1]
	} else {
		left = w[block.Y][block.X-1]
		right = block
	}

	nextPositionX := block.X + addX
	nextPositionY := block.Y + addY

	if nextPositionX < 0 || nextPositionX >= len(w[0]) || nextPositionY < 0 || nextPositionY >= len(w) {
		return w
	}

	nextPosition := w[nextPositionY][nextPositionX]

	if nextPosition.IsEqual(left) || nextPosition.IsEqual(right) {
		w = w.PushBlockRec(nextPosition, block, addX, addY)

		w[nextPositionY][nextPositionX].Type = block.Type
		w[block.Y][block.X].Type = Empty

		return w
	}

	if nextPosition.IsBlock() {
		w = w.PushBlockRec(nextPosition, block, addX, addY)

		if block.IsBlockLeft() {
			if !prev.IsEqual(right) {
				w = w.PushBlockRec(right, block, addX, addY)
			}
		} else {
			if !prev.IsEqual(left) {
				w = w.PushBlockRec(left, block, addX, addY)
			}
		}

		w[nextPositionY][nextPositionX].Type = block.Type
		w[block.Y][block.X].Type = Empty

		return w
	}

	if nextPosition.IsEmpty() {
		if block.IsBlockLeft() {
			if !prev.IsEqual(right) {
				w = w.PushBlockRec(right, block, addX, addY)
			}
		} else {
			if !prev.IsEqual(left) {
				w = w.PushBlockRec(left, block, addX, addY)
			}
		}

		w[nextPositionY][nextPositionX].Type = block.Type
		w[block.Y][block.X].Type = Empty

		return w
	}

	return w
}

func (w WareHouse) PushBlock(block Position, move string) WareHouse {
	robotPosition := w.GetRobotPosition()

	addX := 0
	addY := 0

	switch move {
	case "^":
		addY--
	case ">":
		addX++
	case "v":
		addY++
	case "<":
		addX--
	}

	nextPositionX := block.X + addX
	nextPositionY := block.Y + addY

	for {
		if nextPositionX < 0 || nextPositionX >= len(w[0]) || nextPositionY < 0 || nextPositionY >= len(w) {
			return w
		}

		nextPosition := w[nextPositionY][nextPositionX]

		if nextPosition.IsWall() {
			return w
		} else if nextPosition.IsEmpty() {
			break
		}

		nextPositionX += addX
		nextPositionY += addY
	}

	addX *= -1
	addY *= -1

	for nextPositionX != robotPosition.X || nextPositionY != robotPosition.Y {
		w[nextPositionY][nextPositionX].Type = w[nextPositionY+addY][nextPositionX+addX].Type
		nextPositionX += addX
		nextPositionY += addY
	}

	w[robotPosition.Y][robotPosition.X].Type = Empty

	return w
}

func (w WareHouse) MoveRobot(move string) WareHouse {
	robotPosition := w.GetRobotPosition()

	addX := 0
	addY := 0

	switch move {
	case "^":
		addY--
	case ">":
		addX++
	case "v":
		addY++
	case "<":
		addX--
	}

	nextPositionX := robotPosition.X + addX
	nextPositionY := robotPosition.Y + addY

	if nextPositionX < 0 || nextPositionX >= len(w[0]) || nextPositionY < 0 || nextPositionY >= len(w) {
		return w
	}

	nextPosition := w[nextPositionY][nextPositionX]

	if nextPosition.IsWall() {
		return w
	} else if nextPosition.IsEmpty() {
		w[robotPosition.Y][robotPosition.X].Type = Empty
		w[nextPositionY][nextPositionX].Type = Robot
		return w
	} else if nextPosition.IsBlockRight() || nextPosition.IsBlockLeft() {
		ok := w.CanPushBlockRec(nextPosition, nextPosition, addX, addY)

		if !ok {
			return w
		}

		w = w.PushBlockRec(nextPosition, nextPosition, addX, addY)

		w[nextPositionY][nextPositionX].Type = Robot
		w[robotPosition.Y][robotPosition.X].Type = Empty

		return w
	} else if nextPosition.IsBlock() {
		return w.PushBlock(nextPosition, move)
	}

	return w
}

func parseInput(isWide bool) (WareHouse, []string) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	warehouse := make(WareHouse, 0)
	moves := make([]string, 0)

	parts := strings.Split(string(input), "\n\n")

	y := 0
	warehouseStr := strings.Split(parts[0], "\n")

	for y < len(warehouseStr) {
		x := 0
		row := make([]Position, 0)
		for _, c := range warehouseStr[y] {
			switch c {
			case '#':
				row = append(row, Position{Type: Wall, X: x, Y: y})
				if isWide {
					row = append(row, Position{Type: Wall, X: x + 1, Y: y})
					x++
				}
			case '.':
				row = append(row, Position{Type: Empty, X: x, Y: y})
				if isWide {
					row = append(row, Position{Type: Empty, X: x + 1, Y: y})
					x++
				}
			case 'O':
				if isWide {
					row = append(row, Position{Type: BlockLeft, X: x, Y: y})
					row = append(row, Position{Type: BlockRight, X: x + 1, Y: y})
					x++
				} else {
					row = append(row, Position{Type: Block, X: x, Y: y})
				}
			case '@':
				row = append(row, Position{Type: Robot, X: x, Y: y})
				if isWide {
					row = append(row, Position{Type: Empty, X: x + 1, Y: y})
					x++
				}
			}

			x++
		}

		warehouse = append(warehouse, row)
		y++
	}

	for _, m := range strings.Split(parts[1], "\n") {
		for _, c := range m {
			moves = append(moves, string(c))
		}
	}

	return warehouse, moves
}

func part1() {
	warehouse, moves := parseInput(false)

	for _, m := range moves {
		warehouse = warehouse.MoveRobot(m)
	}

	fmt.Printf("Result part 1: %d\n", warehouse.GetScore())
}

func part2() {
	warehouse, moves := parseInput(true)

	for _, m := range moves {
		warehouse = warehouse.MoveRobot(m)
	}

	fmt.Printf("Result part 2: %d\n", warehouse.GetScore())
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
