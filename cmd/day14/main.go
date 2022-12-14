package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day14"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day14/part1.txt")
	if err != nil {
		panic(err)
	}

	data, err := day14.PrepareData(bytes)
	if err != nil {
		panic(err)
	}

	part1 := day14.Part1(data)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := day14.Part2(data)
	fmt.Printf("Part 2: %d\n", part2)
}
