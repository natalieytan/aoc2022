package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day3"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day3/part1.txt")
	if err != nil {
		panic(err)
	}

	data := day3.PrepareData(bytes)

	part1 := day3.Part1(data)
	fmt.Printf("Part 1: %d\n", part1)

	part2, err := day3.Part2(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2: %d\n", part2)
}
