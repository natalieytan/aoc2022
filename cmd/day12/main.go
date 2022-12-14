package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day12"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day12/part1.txt")
	if err != nil {
		panic(err)
	}

	data := day12.PrepareData(bytes)

	part1, err := day12.Part1(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", part1)

	part2, err := day12.Part2(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2: %d\n", part2)
}
