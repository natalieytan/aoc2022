package day5

func Part1(crateStack CrateStack, instructions []Instruction) string {
	crateStack.moveAll(instructions)
	return crateStack.top()
}

func Part2(crateStack CrateStack, instructions []Instruction) string {
	crateStack.moveAllCrateMover9001(instructions)
	return crateStack.top()
}
