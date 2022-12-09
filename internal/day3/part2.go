package day3

import (
	"errors"
	"unicode"
)

func Part2(elves []string) (int, error) {
	total := 0
	for i := 0; i < len(elves); i = i + 3 {
		commonItems := findCommonItemForElves(elves[i], elves[i+1], elves[i+2])
		if len(commonItems) != 1 {
			return 0, errors.New("found more common items than expected")
		}
		for commonItem := range commonItems {
			total += priorityForItem(commonItem)
		}
	}
	return total, nil
}

func findCommonItemForElves(elf1 string, elf2 string, elf3 string) map[rune]bool {
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

func priorityForItem(item rune) int {
	if unicode.IsUpper(item) {
		return int(item) - 38
	}
	return int(item) - 96
}
