package day3

import (
	"errors"
	"strings"
	"unicode"
)

func PrepareData(input []byte) []string {
	data := string(input)
	return strings.Split(data, "\n")
}

func SumOfCommonPrioritiesForAllRucksacks(rucksacks []string) int {
	total := 0
	for _, rucksack := range rucksacks {
		total += sumOfCommonPrioritiesForRucksack(rucksack)
	}
	return total
}

func sumOfCommonPrioritiesForRucksack(rucksack string) int {
	total := 0
	commonItems := findCommonItems(rucksack)

	for item, _ := range commonItems {
		total += priorityForItem(item)
	}

	return total
}

func findCommonItems(rucksack string) map[rune]bool {
	compartment1 := rucksack[:len(rucksack)/2]
	compartment2 := rucksack[len(rucksack)/2:]

	itemsMap := map[rune]bool{}
	commonItems := map[rune]bool{}

	for _, item := range compartment1 {
		itemsMap[item] = true
	}

	for _, item := range compartment2 {
		if itemsMap[item] {
			commonItems[item] = true
		}
	}
	return commonItems
}

func priorityForItem(item rune) int {
	if unicode.IsUpper(item) {
		return int(item) - 38
	}
	return int(item) - 96
}

func PrioritiesForElfBadges(elves []string) (int, error) {
	total := 0
	for i := 0; i < len(elves); i = i + 3 {
		commonItems := findCommonItemBadgeForElves(elves[i], elves[i+1], elves[i+2])
		if len(commonItems) != 1 {
			return 0, errors.New("found more common items than expected")
		}
		for commonItem := range commonItems {
			total += priorityForItem(commonItem)
		}
	}
	return total, nil
}

func findCommonItemBadgeForElves(elf1 string, elf2 string, elf3 string) map[rune]bool {
	elf1Map := map[rune]bool{}
	elf2Map := map[rune]bool{}
	elf3Map := map[rune]bool{}

	for _, item := range elf1 {
		elf1Map[item] = true
	}

	for _, item := range elf2 {
		if elf1Map[item] {
			elf2Map[item] = true
		}
	}

	for _, item := range elf3 {
		if elf2Map[item] {
			elf3Map[item] = true
		}
	}

	return elf3Map
}
