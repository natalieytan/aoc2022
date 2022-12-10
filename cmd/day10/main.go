package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day10"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day10/part1.txt")
	if err != nil {
		panic(err)
	}

	data, err := day10.PrepareData(bytes)
	if err != nil {
		panic(err)
	}

	part1 := day10.Part1(data)
	fmt.Printf("Part 1: %d\n", part1)

	fmt.Printf("Part 2:\n")
	day10.PrintPart2(data)
}
