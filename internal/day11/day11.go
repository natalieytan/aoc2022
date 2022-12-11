package day11

import (
	"sort"
)

func Part1(monkeys []*Monkey) int {
	worryModifier := func(item int) int {
		return item / 3
	}
	return calculateMonkeyBusinessAfterRounds(20, monkeys, worryModifier)
}

func Part2(monkeys []*Monkey) int {
	worryModifier := func(item int) int {
		return item % leastCommonMultipleOfMonkeyDivisors(monkeys)
	}
	return calculateMonkeyBusinessAfterRounds(10000, monkeys, worryModifier)
}

func calculateMonkeyBusinessAfterRounds(rounds int, monkeys []*Monkey, worryModifier func(int) int) int {
	for round := 1; round <= rounds; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				itemNewWorryLevel := worryModifier(monkey.worryLevelAfterInspectingItem(item))
				monkey.inspections += 1
				monkeyThrowTo := monkeys[monkey.itemToBeThrownToMonkeyNo(itemNewWorryLevel)]
				monkeyThrowTo.items = append(monkeyThrowTo.items, itemNewWorryLevel)
				monkey.items = monkey.items[1:]
			}
		}
	}

	return calculateMonkeyBusiness(monkeys)
}

func calculateMonkeyBusiness(monkeys []*Monkey) int {
	monkeyInspections := []int{}
	for _, monkey := range monkeys {
		monkeyInspections = append(monkeyInspections, monkey.inspections)
	}
	sort.Ints(monkeyInspections)
	top2 := monkeyInspections[len(monkeyInspections)-2:]

	return top2[0] * top2[1]
}
