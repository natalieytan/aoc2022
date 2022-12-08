package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day8"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day8/part1.txt")
	if err != nil {
		panic(err)
	}

	data, err := day8.PrepareData(bytes)
	if err != nil {
		panic(err)
	}

	part1 := day8.VisibleTreesOutsideGrid(data)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := day8.MaxScenicScore(data)
	fmt.Printf("Part 2: %d\n", part2)
}
