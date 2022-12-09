package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day5"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day5/part1.txt")
	if err != nil {
		panic(err)
	}

	instructionData, err := day5.PrepareData(bytes)
	if err != nil {
		panic(err)
	}

	part1 := day5.Part1(day5.NewCrateStack(), instructionData)
	fmt.Printf("Part 1: %s\n", part1)

	part2 := day5.Part2(day5.NewCrateStack(), instructionData)
	fmt.Printf("Part 2: %s\n", part2)
}
