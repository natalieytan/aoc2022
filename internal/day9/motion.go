package day9

import (
	"errors"
	"strconv"
	"strings"
)

type Direction int

const (
	UnknownDirection Direction = iota
	Up
	Down
	Left
	Right
)

type Motion struct {
	Direction Direction
	Distance  int
}

func PrepareData(input []byte) ([]Motion, error) {
	data := string(input)
	rawMotions := strings.Split(data, "\n")
	motions := make([]Motion, len(rawMotions))
	for i, rawMotion := range rawMotions {
		motion, err := parseMotion(rawMotion)
		if err != nil {
			return motions, err
		}
		motions[i] = motion
	}

	return motions, nil
}

func parseMotion(rawMotion string) (Motion, error) {
	directionAndDistance := strings.Split(rawMotion, " ")
	if len(directionAndDistance) != 2 {
		return Motion{}, errors.New("error parsing motion")
	}
	distance, err := strconv.Atoi(directionAndDistance[1])
	if err != nil {
		return Motion{}, nil
	}
	direction, err := parseDirection(directionAndDistance[0])
	if err != nil {
		return Motion{}, nil
	}

	return Motion{
		Direction: direction,
		Distance:  distance,
	}, nil

}

func parseDirection(rawDirection string) (Direction, error) {
	switch rawDirection {
	case "R":
		return Right, nil
	case "U":
		return Up, nil
	case "L":
		return Left, nil
	case "D":
		return Down, nil
	default:
		return UnknownDirection, errors.New("invalid direction")
	}
}
