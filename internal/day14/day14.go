package day14

func Part1(gridData GridData) int {
	sandUnits := 0
	rg := newGrid(gridData)
	droppable := true

	for droppable {
		sand := newSandStartCoordinate()
		_, ok := rg.dropUntilStop(sand)
		if !ok {
			droppable = false
		} else {
			sandUnits += 1
		}
	}

	return sandUnits
}

func Part2(gridData GridData) int {
	sandUnits := 0
	rg := newGrid(gridData)
	droppable := true

	for droppable {
		sand := newSandStartCoordinate()
		_, ok := rg.dropUntilStopWithFloor(sand)
		if !ok {
			droppable = false
		} else {
			sandUnits += 1
		}
	}

	return sandUnits
}
