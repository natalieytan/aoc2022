package day14

type GridData struct {
	occupiedCoordinates CoordinateMap
	lowestY             int
}

type Grid struct {
	occupiedCoordinates CoordinateMap
	lowestY             int
}

func newGrid(gd GridData) Grid {
	return Grid{
		occupiedCoordinates: gd.occupiedCoordinates.copy(),
		lowestY:             gd.lowestY,
	}
}

func (g *Grid) dropUntilStop(sand Coordinate) (Coordinate, bool) {
	if sand.y > g.lowestY {
		return Coordinate{}, false
	}

	downCoordinate, ok := g.tryDropDown(sand)
	if ok {
		return g.dropUntilStop(downCoordinate)
	}
	leftCoordinate, ok := g.tryDropLeft(sand)
	if ok {
		return g.dropUntilStop(leftCoordinate)
	}
	rightCoordinate, ok := g.tryDropRight(sand)
	if ok {
		return g.dropUntilStop(rightCoordinate)
	}

	if g.occupiedCoordinates[sand.str()] {
		return Coordinate{}, false
	}

	g.occupiedCoordinates[sand.str()] = true

	return sand, true
}

func (g *Grid) dropUntilStopWithFloor(sand Coordinate) (Coordinate, bool) {
	downCoordinate, ok := g.tryDropDownWithFloor(sand)
	if ok {
		return g.dropUntilStopWithFloor(downCoordinate)
	}
	leftCoordinate, ok := g.tryDropLeftWithFloor(sand)
	if ok {
		return g.dropUntilStopWithFloor(leftCoordinate)
	}
	rightCoordinate, ok := g.tryDropRightWithFloor(sand)
	if ok {
		return g.dropUntilStopWithFloor(rightCoordinate)
	}

	if g.occupiedCoordinates[sand.str()] {
		return Coordinate{}, false
	}

	g.occupiedCoordinates[sand.str()] = true
	return sand, true
}

func (g *Grid) tryDropDownWithFloor(sand Coordinate) (Coordinate, bool) {
	downCoordinate := sand.down()
	if !g.occupiedCoordinates[downCoordinate.str()] && downCoordinate.y < g.lowestY+2 {
		return downCoordinate, true
	}
	return downCoordinate, false
}

func (g *Grid) tryDropLeftWithFloor(sand Coordinate) (Coordinate, bool) {
	downCoordinate := sand.diagonalLeft()
	if !g.occupiedCoordinates[downCoordinate.str()] && downCoordinate.y < g.lowestY+2 {
		return downCoordinate, true
	}
	return downCoordinate, false
}

func (g *Grid) tryDropRightWithFloor(sand Coordinate) (Coordinate, bool) {
	downCoordinate := sand.diagonalRight()
	if !g.occupiedCoordinates[downCoordinate.str()] && downCoordinate.y < g.lowestY+2 {
		return downCoordinate, true
	}
	return downCoordinate, false
}

func (g *Grid) tryDropDown(sand Coordinate) (Coordinate, bool) {
	downCoordinate := sand.down()
	occupied := g.occupiedCoordinates[downCoordinate.str()]
	if !occupied {
		return downCoordinate, true
	}
	return downCoordinate, false
}

func (g *Grid) tryDropLeft(sand Coordinate) (Coordinate, bool) {
	downCoordinate := sand.diagonalLeft()
	if !g.occupiedCoordinates[downCoordinate.str()] {
		return downCoordinate, true
	}
	return downCoordinate, false
}

func (g *Grid) tryDropRight(sand Coordinate) (Coordinate, bool) {
	downCoordinate := sand.diagonalRight()
	if !g.occupiedCoordinates[downCoordinate.str()] {
		return downCoordinate, true
	}
	return downCoordinate, false
}
