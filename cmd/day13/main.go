package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day13"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day13/part1.txt")
	if err != nil {
		panic(err)
	}

	data, err := day13.PrepareData(bytes)
	if err != nil {
		panic(err)
	}

	part1 := day13.Part1(data)
	fmt.Printf("Part 1: %d\n", part1)

	data2, err := day13.PrepareData2(bytes)
	if err != nil {
		panic(err)
	}

	part2 := day13.Part2(data2)
	fmt.Printf("Part 2: %d\n", part2)
}
