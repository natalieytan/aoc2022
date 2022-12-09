package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day7"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day7/part1.txt")
	if err != nil {
		panic(err)
	}

	data, err := day7.PrepareData(bytes)
	if err != nil {
		panic(err)
	}

	rootDirectory := day7.BuildRootDirectory(data)

	part1 := day7.Part1(rootDirectory)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := day7.Part2(rootDirectory)
	fmt.Printf("Part 2: %d\n", part2)
}
