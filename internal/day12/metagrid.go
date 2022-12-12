package day12

import "errors"

type MetaGrid struct {
	rows    int
	cols    int
	start   Coordinate
	end     Coordinate
	heights [][]int
}

type Coordinate struct {
	row int
	col int
}

func newMetaGrid(grid []string) MetaGrid {
	gridHeight := len(grid)
	gridWidth := len(grid[0])
	startPosition := Coordinate{}
	endPosition := Coordinate{}
	heights := make([][]int, len(grid))

	for rowIdx, row := range grid {
		newRow := make([]int, len(row))
		for colIdx, square := range row {
			if square == 'E' {
				newRow[colIdx] = 26
				endPosition = Coordinate{
					row: rowIdx,
					col: colIdx,
				}
			} else if square == 'S' {
				startPosition = Coordinate{
					row: rowIdx,
					col: colIdx,
				}
				newRow[colIdx] = 1
			} else {
				newRow[colIdx] = int(square) - 96
			}
		}
		heights[rowIdx] = newRow
	}

	return MetaGrid{
		rows:    gridHeight,
		cols:    gridWidth,
		start:   startPosition,
		end:     endPosition,
		heights: heights,
	}
}

func (m MetaGrid) shortestDistanceToEndFromCoordinates(startingCoordinates []Coordinate) (int, error) {
	shortest, _ := m.shortestDistanceToEndFromCoordinate(m.start)

	for _, coordinate := range startingCoordinates {
		distance, err := m.shortestDistanceToEndFromCoordinate(coordinate)
		if err == nil && (distance < shortest || shortest == 0) {
			shortest = distance
		}
	}

	if shortest == 0 {
		return 0, errors.New("unable to find path to end")
	}

	return shortest, nil
}

func (m MetaGrid) shortestDistanceToEndFromCoordinate(startingCoordinate Coordinate) (int, error) {
	distanceGrid := m.sameSizeEmptyGrid()
	searchSurroundingCoords := []Coordinate{startingCoordinate}

	for len(searchSurroundingCoords) > 0 {
		coord := searchSurroundingCoords[0]
		adjacentCoords := m.adjacentCoordinates(coord)
		for _, adjacentCoord := range adjacentCoords {
			currentHeight := m.heights[coord.row][coord.col]
			if currentHeight+1 >= m.heights[adjacentCoord.row][adjacentCoord.col] {
				currentDistance := distanceGrid[coord.row][coord.col] + 1
				if adjacentCoord.col == m.end.col && adjacentCoord.row == m.end.row {
					return currentDistance, nil
				}
				maxDistance := distanceGrid[adjacentCoord.row][adjacentCoord.col]
				if (maxDistance == 0) || currentDistance < maxDistance {
					if !(startingCoordinate.row == adjacentCoord.row && startingCoordinate.col == adjacentCoord.col) {
						distanceGrid[adjacentCoord.row][adjacentCoord.col] = currentDistance
						searchSurroundingCoords = append(searchSurroundingCoords, adjacentCoord)
					}
				}
			}
		}
		searchSurroundingCoords = searchSurroundingCoords[1:]
	}

	return 0, errors.New("unable to find path to end")
}

func (m MetaGrid) coordinatesWithLowestElevation() []Coordinate {
	coordinates := []Coordinate{}
	for rowIdx, row := range m.heights {
		for colIdx, height := range row {
			if height == 1 {
				coordinates = append(coordinates, Coordinate{
					row: rowIdx,
					col: colIdx,
				})
			}
		}
	}
	return coordinates
}

func (m MetaGrid) adjacentCoordinates(c Coordinate) []Coordinate {
	coordinates := []Coordinate{}

	if c.row != 0 {
		coordinates = append(coordinates, Coordinate{
			row: c.row - 1,
			col: c.col,
		})
	}

	if c.row != m.rows-1 {
		coordinates = append(coordinates, Coordinate{
			row: c.row + 1,
			col: c.col,
		})
	}

	if c.col != 0 {
		coordinates = append(coordinates, Coordinate{
			row: c.row,
			col: c.col - 1,
		})
	}

	if c.col != m.cols-1 {
		coordinates = append(coordinates, Coordinate{
			row: c.row,
			col: c.col + 1,
		})
	}

	return coordinates
}

func (m MetaGrid) sameSizeEmptyGrid() [][]int {
	grid := make([][]int, m.rows)
	for i := 0; i < m.rows; i++ {
		grid[i] = make([]int, m.cols)
	}

	return grid
}
