package day11

func leastCommonMultipleOfMonkeyDivisors(monkeys []*Monkey) int {
	divisors := []int{}
	for _, monkey := range monkeys {
		divisors = append(divisors, monkey.decision.divisbleBy)
	}
	return leastCommonMultiple(divisors[0], divisors[1], divisors[2:]...)
}

func leastCommonMultiple(val1 int, val2 int, vals ...int) int {
	lcm := val1 * val2 / greatestCommonDivisor(val1, val2)
	for i := 0; i < len(vals); i++ {
		lcm = leastCommonMultiple(lcm, vals[i])
	}

	return lcm
}

func greatestCommonDivisor(val1, val2 int) int {
	for val2 != 0 {
		temp := val2
		val2 = val1 % val2
		val1 = temp
	}
	return val1
}
