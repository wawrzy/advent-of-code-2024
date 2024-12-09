package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	ID          int
	Size        int
	IsFreeSpace bool
}

func parseInput() []Block {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	blocks := make([]Block, 0, len(string(input)))

	id := 0

	for i, sizeStr := range strings.Split(string(input), "") {
		size, err := strconv.Atoi(sizeStr)

		if i%2 == 0 {
			if err != nil {
				panic(err)
			}

			blocks = append(blocks, Block{ID: id, Size: size})
			id++
		} else {
			blocks = append(blocks, Block{Size: size, IsFreeSpace: true})
		}
	}

	return blocks
}

func part1() {
	disk := parseInput()

	decompressed := make([]int, 0)

	for _, block := range disk {
		for i := 0; i < block.Size; i++ {
			if block.IsFreeSpace {
				decompressed = append(decompressed, -1)
			} else {
				decompressed = append(decompressed, block.ID)
			}
		}
	}

	compacted := make([]int, 0)

	freeSpaceIdx := 0
	dataIdx := len(decompressed) - 1

	for dataIdx >= 0 && freeSpaceIdx < len(decompressed) && dataIdx >= freeSpaceIdx {
		if decompressed[dataIdx] == -1 {
			dataIdx--
			continue
		}

		if decompressed[freeSpaceIdx] != -1 {
			compacted = append(compacted, decompressed[freeSpaceIdx])
			freeSpaceIdx++
		} else {
			compacted = append(compacted, decompressed[dataIdx])
			dataIdx--
			freeSpaceIdx++
		}
	}

	res := 0

	for i, v := range compacted {
		if v != -1 {
			res += i * v
		}
	}

	fmt.Printf("Result part 1: %d\n", res)
}

func part2() {
	disk := parseInput()

	endIdx := len(disk) - 1

	for endIdx >= 0 {
		if disk[endIdx].IsFreeSpace {
			endIdx--
			continue
		}

		freeSpaceIdx := 0
		for freeSpaceIdx < endIdx {
			if !disk[freeSpaceIdx].IsFreeSpace {
				freeSpaceIdx++
				continue
			}

			if disk[freeSpaceIdx].Size >= disk[endIdx].Size {
				break
			}

			freeSpaceIdx++
		}

		if freeSpaceIdx != endIdx {
			fileBlock := disk[endIdx]

			disk[endIdx].ID = -1
			disk[endIdx].IsFreeSpace = true

			disk[freeSpaceIdx].Size -= fileBlock.Size

			disk = slices.Insert(disk, freeSpaceIdx, fileBlock)
			endIdx--
		} else {
			endIdx--
		}
	}

	res := 0
	i := 0

	for _, block := range disk {
		if block.IsFreeSpace {
			i += block.Size
			continue
		} else {
			for j := 0; j < block.Size; j++ {
				res += i * block.ID
				i++
			}
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
