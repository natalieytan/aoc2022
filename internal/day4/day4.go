package day4

func Part1(assignmentPairs []AssignmentPair) int {
	total := 0
	for _, assignmentPair := range assignmentPairs {
		if assignmentPair.fullyContains() {
			total += 1
		}
	}

	return total
}

func Part2(assignmentPairs []AssignmentPair) int {
	total := 0
	for _, assignmentPair := range assignmentPairs {
		if assignmentPair.overlaps() {
			total += 1
		}
	}

	return total
}
