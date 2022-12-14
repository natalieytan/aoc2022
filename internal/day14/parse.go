package day14

import (
	"errors"
	"strconv"
	"strings"
)

func PrepareData(input []byte) (GridData, error) {
	rockCoordinateMap := map[string]bool{}
	data := string(input)
	rawRockPaths := strings.Split(data, "\n")
	lowestY := 0

	for _, rawRockPath := range rawRockPaths {
		rockPathCoordinates, err := parseRockPath(rawRockPath)
		if err != nil {
			return GridData{}, err
		}
		for i := 0; i < len(rockPathCoordinates)-1; i++ {
			rockCoordinates := rockPathCoordinates[i].coordsBetween(rockPathCoordinates[i+1])
			for _, rockCoordinate := range rockCoordinates {
				rockCoordinateMap[rockCoordinate.str()] = true
				if rockCoordinate.y > lowestY {
					lowestY = rockCoordinate.y
				}
			}
		}
	}

	return GridData{
		occupiedCoordinates: rockCoordinateMap,
		lowestY:             lowestY,
	}, nil
}

func parseRockPath(rawRockPath string) ([]Coordinate, error) {
	rawRockCoordinates := strings.Split(rawRockPath, " -> ")
	rockCoordinates := make([]Coordinate, len(rawRockCoordinates))

	for i, rawRockCoordiante := range rawRockCoordinates {
		rockCoordinate, err := parseCoordinate(rawRockCoordiante)
		if err != nil {
			return rockCoordinates, err
		}
		rockCoordinates[i] = rockCoordinate

	}
	return rockCoordinates, nil
}

func parseCoordinate(rawCoord string) (Coordinate, error) {
	coords := strings.Split(rawCoord, ",")
	if len(coords) != 2 {
		return Coordinate{}, errors.New("error parsing coordinate")
	}
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		return Coordinate{}, err
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return Coordinate{}, err
	}

	return Coordinate{
		x: x,
		y: y,
	}, err
}
