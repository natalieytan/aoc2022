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

	data, err := day2.PrepareDataPt1(bytes)
	if err != nil {
		panic(err)
	}
	totalScore := day2.TotalScoreForPlayerPt1(data)
	fmt.Printf("Part 1: %d\n", totalScore)

	dataPt2, err := day2.PrepareDataPt2(bytes)
	if err != nil {
		panic(err)
	}
	totalScorePt2 := day2.TotalScoreForPlayerPt2(dataPt2)
	fmt.Printf("Part 1: %d\n", totalScorePt2)
}
