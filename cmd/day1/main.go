package main

import (
	"fmt"
	"os"

	"github.com/natalieytan/aoc2022/internal/day1"
)

func main() {
	bytes, err := os.ReadFile("./inputs/day1/part1.txt")
	if err != nil {
		panic(err)
	}
	caloriesData, err := day1.PrepareCaloriesDataPerElf(bytes)
	if err != nil {
		panic(err)
	}

	mostCalories := day1.MaxCaloriesForAllElves(caloriesData)
	fmt.Printf("Part 1: %d\n", mostCalories)

	top3 := day1.SumOfTop3ElfCalories(caloriesData)
	fmt.Printf("Part 2: %d\n", top3)

}
