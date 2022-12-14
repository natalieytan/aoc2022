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

	droppedCoordinate, ok := g.tryDrop(sand)
	if ok {
		return g.dropUntilStop(droppedCoordinate)
	}

	if g.occupiedCoordinates[sand.str()] {
		return Coordinate{}, false
	}

	g.occupiedCoordinates[sand.str()] = true
	return sand, true
}

func (g *Grid) tryDrop(sand Coordinate) (Coordinate, bool) {
	// try drop down
	downCoordinate := sand.down()
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

func (g *Grid) dropUntilStopWithFloor(sand Coordinate) (Coordinate, bool) {
	droppedCordinate, ok := g.tryDropWithFloor(sand)
	if ok {
		return g.dropUntilStopWithFloor(droppedCordinate)
	}

	if g.occupiedCoordinates[sand.str()] {
		return Coordinate{}, false
	}

	g.occupiedCoordinates[sand.str()] = true
	return sand, true
}

func (g *Grid) tryDropWithFloor(sand Coordinate) (Coordinate, bool) {
	// try drop down
	downCoordinate := sand.down()
	if !g.occupiedCoordinates[downCoordinate.str()] && downCoordinate.y < g.floor() {
		return downCoordinate, true
	}
	// followed by diagonally left
	leftCoordinate := sand.diagonalLeft()
	if !g.occupiedCoordinates[leftCoordinate.str()] && downCoordinate.y < g.floor() {
		return leftCoordinate, true
	}
	// followed by diagonally right
	rightCoordinate := sand.diagonalRight()
	if !g.occupiedCoordinates[rightCoordinate.str()] && downCoordinate.y < g.floor() {
		return rightCoordinate, true
	}
	// not droppable - sand has stopped
	return Coordinate{}, false
}

func (g *Grid) floor() int {
	return g.lowestY + 2
}
