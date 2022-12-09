package day1

import (
	"sort"
)

func Part1(allCaloriesPerElf [][]int) int {
	maxCalories := 0
	for _, totalCalories := range totalCaloriesPerElf(allCaloriesPerElf) {
		if totalCalories > maxCalories {
			maxCalories = totalCalories
		}
	}
	return maxCalories
}

func Part2(allCaloriesPerElf [][]int) int {
	total := totalCaloriesPerElf(allCaloriesPerElf)
	sort.Ints(total)
	top3 := total[len(allCaloriesPerElf)-3:]
	sumOfTop3 := 0
	for _, n := range top3 {
		sumOfTop3 += n
	}
	return sumOfTop3
}

/*
Takes the caloric data for all elves and returns the total calories for each elf.
*/
func totalCaloriesPerElf(allCaloriesPerElf [][]int) []int {
	totalCalories := make([]int, len(allCaloriesPerElf))
	for i, calory := range allCaloriesPerElf {
		total := 0
		for _, c := range calory {
			total += c
		}
		totalCalories[i] = total
	}
	return totalCalories
}
