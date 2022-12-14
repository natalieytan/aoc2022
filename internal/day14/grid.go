package day14

type GridData struct {
	occupiedCoordinates CoordinateMap
	lowestY             int
}

type Grid struct {
	occupiedCoordinates CoordinateMap
	lowestY             int
	hasFloor            bool
}

func newGrid(gd GridData, hasFloor bool) Grid {
	return Grid{
		occupiedCoordinates: gd.occupiedCoordinates.copy(),
		lowestY:             gd.lowestY,
		hasFloor:            hasFloor,
	}
}

func (g *Grid) countDroppableSand() int {
	sandUnits := 0
	droppable := true

	for droppable {
		sand := newSandStartCoordinate()
		_, ok := g.dropUntilStop(sand)
		if !ok {
			droppable = false
		} else {
			sandUnits += 1
		}
	}

	return sandUnits
}

func (g *Grid) dropUntilStop(sand Coordinate) (Coordinate, bool) {
	if !g.hasFloor && sand.y > g.lowestY {
		// no floor. overflow. no longer droppable
		return Coordinate{}, false
	}

	droppedCoordinate, ok := g.tryDrop(sand)
	if ok {
		// sand can be dropped - keep dropping
		return g.dropUntilStop(droppedCoordinate)
	}

	// sand can no longer flow - no longer droppable
	if g.occupiedCoordinates[sand.str()] {
		return Coordinate{}, false
	}

	// sand has stopped and taken up coordinate
	g.occupiedCoordinates[sand.str()] = true
	return sand, true
}

func (g *Grid) tryDrop(sand Coordinate) (Coordinate, bool) {
	// try drop down
	downCoordinate := sand.down()
	if downCoordinate.y >= g.floor() && g.hasFloor {
		// not droppable - sand has hit the floor
		return Coordinate{}, false
	}
	if !g.occupiedCoordinates[downCoordinate.str()] {
		return downCoordinate, true
	}
	// followed by diagonally left
	leftCoordinate := sand.diagonalLeft()
	if !g.occupiedCoordinates[leftCoordinate.str()] {
		return leftCoordinate, true
	}
	// followed by diagonally right
	rightCoordinate := sand.diagonalRight()
	if !g.occupiedCoordinates[rightCoordinate.str()] {
		return rightCoordinate, true
	}
	// not droppable - sand has stopped
	return Coordinate{}, false
}

func (g *Grid) floor() int {
	return g.lowestY + 2
}
