package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day4"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day4/part1.txt")
	if err != nil {
		panic(err)
	}

	data, err := day4.PrepareData(bytes)

	if err != nil {
		panic(err)
	}

	part1 := day4.Part1(data)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := day4.Part2(data)
	fmt.Printf("Part 2: %d\n", part2)
}
