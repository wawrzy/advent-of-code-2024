package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strconv"
	"strings"
	"time"
)

type Robots struct {
	positionX int
	positionY int
	velocityX int
	velocityY int
}

func parseInput() []Robots {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	robots := make([]Robots, 0)

	for _, p := range strings.Split(string(input), "\n") {
		parts := strings.Fields(p)

		pos := strings.FieldsFunc(parts[0], func(r rune) bool {
			return r == '='
		})
		position := strings.FieldsFunc(pos[1], func(r rune) bool {
			return r == ','
		})

		vel := strings.FieldsFunc(parts[1], func(r rune) bool {
			return r == '='
		})
		velocity := strings.FieldsFunc(vel[1], func(r rune) bool {
			return r == ','
		})

		positionX, _ := strconv.Atoi(position[0])
		positionY, _ := strconv.Atoi(position[1])
		velocityX, _ := strconv.Atoi(velocity[0])
		velocityY, _ := strconv.Atoi(velocity[1])

		robots = append(robots, Robots{
			positionX: positionX,
			positionY: positionY,
			velocityX: velocityX,
			velocityY: velocityY,
		})
	}

	return robots
}

func display(spaces [][]int, removeMid bool) {
	// Clear the screen
	fmt.Print("\033[H\033[2J")

	for i, s := range spaces {
		if (len(spaces)-1)/2 == i && removeMid {
			fmt.Println()
			continue
		}

		for j, ss := range s {
			if (len(s)-1)/2 == j && removeMid {
				fmt.Print(" ")
				continue
			}

			if ss == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", ss)
			}
		}
		fmt.Println()
	}
}

func countRobotsInQuadrants(spaces [][]int) int {
	var quadrants [4]int

	for i, s := range spaces {
		if (len(spaces)-1)/2 == i {
			continue
		}

		for j, ss := range s {
			if ss == 0 {
				continue
			}
			if (len(s)-1)/2 == j {
				continue
			}

			if i < len(spaces)/2 && j < len(s)/2 {
				quadrants[0] += spaces[i][j]
			} else if i < len(spaces)/2 && j >= len(s)/2 {
				quadrants[1] += spaces[i][j]
			} else if i >= len(spaces)/2 && j < len(s)/2 {
				quadrants[2] += spaces[i][j]
			} else {
				quadrants[3] += spaces[i][j]
			}
		}
	}

	res := 1
	for _, q := range quadrants {
		res *= q
	}
	return res
}

func displayPng(spaces [][]int, frame int) {
	zoom := 10

	tall := 103 * zoom
	wide := 101 * zoom

	new_png_file := fmt.Sprintf("frame_%d.png", frame)

	myblack := color.RGBA{0, 0, 0, 255}
	mywhite := color.RGBA{255, 255, 255, 255}

	myimage := image.NewRGBA(image.Rect(0, 0, wide, tall))
	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{myblack}, image.ZP, draw.Src)

	for y, s := range spaces {
		for x, ss := range s {
			red_rect := image.Rect(x*zoom, y*zoom, (x+1)*zoom, (y+1)*zoom)

			if ss != 0 {
				draw.Draw(myimage, red_rect, &image.Uniform{mywhite}, image.ZP, draw.Src)
			}
		}
	}

	myfile, err := os.Create(new_png_file)
	if err != nil {
		panic(err)
	}
	defer myfile.Close()
	png.Encode(myfile, myimage)
}

func part1() {
	robots := parseInput()

	tall := 103
	wide := 101

	spaces := make([][]int, 0, tall)
	for i := 0; i < tall; i++ {
		spaces = append(spaces, make([]int, 0, wide))
		for j := 0; j < wide; j++ {
			spaces[i] = append(spaces[i], 0)
		}
	}

	// Initialize the spaces
	for _, r := range robots {
		spaces[r.positionY][r.positionX] += 1
	}

	maxSeconds := 100

	for i := 0; i < maxSeconds; i++ {
		for j, r := range robots {
			nextPositionX := r.positionX + r.velocityX
			nextPositionY := r.positionY + r.velocityY

			if nextPositionX < 0 {
				nextPositionX += wide
			}
			if nextPositionX >= wide {
				nextPositionX -= wide
			}
			if nextPositionY < 0 {
				nextPositionY += tall
			}
			if nextPositionY >= tall {
				nextPositionY -= tall
			}

			spaces[r.positionY][r.positionX] -= 1
			spaces[nextPositionY][nextPositionX] += 1

			robots[j].positionX = nextPositionX
			robots[j].positionY = nextPositionY
		}
	}

	res := countRobotsInQuadrants(spaces)

	fmt.Printf("Result part 1: %d\n", res)
}

func part2() {
	robots := parseInput()

	tall := 103
	wide := 101

	spaces := make([][]int, 0, tall)
	for i := 0; i < tall; i++ {
		spaces = append(spaces, make([]int, 0, wide))
		for j := 0; j < wide; j++ {
			spaces[i] = append(spaces[i], 0)
		}
	}

	// Initialize the spaces
	for _, r := range robots {
		spaces[r.positionY][r.positionX] += 1
	}

	maxSeconds := 10000

	for i := 0; i < maxSeconds; i++ {
		for j, r := range robots {
			nextPositionX := r.positionX + r.velocityX
			nextPositionY := r.positionY + r.velocityY

			if nextPositionX < 0 {
				nextPositionX += wide
			}
			if nextPositionX >= wide {
				nextPositionX -= wide
			}
			if nextPositionY < 0 {
				nextPositionY += tall
			}
			if nextPositionY >= tall {
				nextPositionY -= tall
			}

			spaces[r.positionY][r.positionX] -= 1
			spaces[nextPositionY][nextPositionX] += 1

			robots[j].positionX = nextPositionX
			robots[j].positionY = nextPositionY
		}

		displayPng(spaces, i)
	}
}

func main() {
	now := time.Now()
	part1()
	duration := time.Since(now)

	fmt.Printf("Duration part 1: %v\n", duration)

	part2()
}
