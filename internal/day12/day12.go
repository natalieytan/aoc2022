package day12

func Part1(grid []string) (int, error) {
	metaGrid := newMetaGrid(grid)
	return metaGrid.shortestDistanceToEndFromCoordinate(metaGrid.start)
}

func Part2(grid []string) (int, error) {
	metaGrid := newMetaGrid(grid)
	return metaGrid.shortestDistanceToEndFromCoordinates(metaGrid.coordinatesWithLowestElevation())
}
