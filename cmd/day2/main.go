package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day2"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day2/part1.txt")
	if err != nil {
		panic(err)
	}

	data, err := day2.PrepareData(bytes)
	if err != nil {
		panic(err)
	}

	dataPt1, err := day2.CovertToSymbolMatchups(data)
	if err != nil {
		panic(err)
	}
	totalScore := day2.Part1(dataPt1)
	fmt.Printf("Part 1: %d\n", totalScore)

	dataPt2, err := day2.CovertToOutComeMatchups(data)
	if err != nil {
		panic(err)
	}
	totalScorePt2 := day2.Part2(dataPt2)
	fmt.Printf("Part 1: %d\n", totalScorePt2)
}
