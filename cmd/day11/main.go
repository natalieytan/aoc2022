package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day11"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day11/part1.txt")
	if err != nil {
		panic(err)
	}

	data, err := day11.PrepareData(bytes)
	if err != nil {
		panic(err)
	}

	part1 := day11.Part1(data)
	fmt.Printf("Part 1: %d\n", part1)

	data2, err := day11.PrepareData(bytes)
	if err != nil {
		panic(err)
	}

	part2 := day11.Part2(data2)
	fmt.Printf("Part 2: %d\n", part2)
}
