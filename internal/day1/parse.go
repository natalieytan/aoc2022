package day1

import (
	"strconv"
	"strings"
)

/*
Parses caloric information to an array containing array of integers.
Within the outer array, each array represents the caloric information of the food for on elf.
Witihin the inner array, each element (int) represents the caloric information of a food iteam.
*/
func PrepareData(input []byte) ([][]int, error) {
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
