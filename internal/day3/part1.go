package day3

func Part1(rucksacks []string) int {
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
