package day1

import (
	"sort"
	"strconv"
	"strings"
)

/*
Parses caloric information to an array containing array of integers.
Within the outer array, each array represents the caloric information of the food for on elf.
Witihin the inner array, each element (int) represents the caloric information of a food iteam.
*/
func PrepareCaloriesDataPerElf(input []byte) ([][]int, error) {
	data := string(input)
	rawCaloriesForElves := strings.Split(data, "\n\n")
	caloricDataForElves := make([][]int, len(rawCaloriesForElves))

	for i, rawCaloriesForElf := range rawCaloriesForElves {
		rawCalories := strings.Split(rawCaloriesForElf, "\n")
		caloricDataForElf := make([]int, len(rawCalories))
		for _, rawCalory := range rawCalories {
			caloricData, err := strconv.Atoi(rawCalory)
			if err != nil {
				return caloricDataForElves, err
			} else {
				caloricDataForElf = append(caloricDataForElf, caloricData)
			}
		}
		caloricDataForElves[i] = caloricDataForElf
	}
	return caloricDataForElves, nil
}

/*
Takes the caloric data for all elves and returns the total calories for each elf.
*/
func TotalCaloriesPerElf(allCaloriesPerElf [][]int) []int {
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

func SumOfTop3ElfCalories(allCaloriesPerElf [][]int) int {
	total := TotalCaloriesPerElf(allCaloriesPerElf)
	sort.Ints(total)
	top3 := total[len(allCaloriesPerElf)-3:]
	sumOfTop3 := 0
	for _, n := range top3 {
		sumOfTop3 += n
	}
	return sumOfTop3
}

func MaxCaloriesForAllElves(allCaloriesPerElf [][]int) int {
	maxCalories := 0
	for _, totalCalories := range TotalCaloriesPerElf(allCaloriesPerElf) {
		if totalCalories > maxCalories {
			maxCalories = totalCalories
		}
	}
	return maxCalories
}
