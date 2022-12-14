package day14

func Part1(gridData GridData) int {
	rg := newGrid(gridData, false)
	return rg.countDroppableSand()

}

func Part2(gridData GridData) int {
	rg := newGrid(gridData, true)
	return rg.countDroppableSand()
}
