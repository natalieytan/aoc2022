package day14

import "fmt"

type CoordinateMap map[string]bool

func (c CoordinateMap) copy() CoordinateMap {
	newMap := map[string]bool{}
	for k, v := range c {
		newMap[k] = v
	}
	return newMap
}

type Coordinate struct {
	x int
	y int
}

func newSandStartCoordinate() Coordinate {
	return Coordinate{
		x: 500,
		y: 0,
	}
}

func (c Coordinate) str() string {
	return fmt.Sprintf("%d-%d", c.x, c.y)
}

func (c Coordinate) coordsBetween(c2 Coordinate) []Coordinate {
	coords := []Coordinate{}
	if c.x == c2.x {
		if c.y > c2.y {
			for i := c2.y; i <= c.y; i++ {
				coords = append(coords, Coordinate{x: c.x, y: i})
			}
		} else {
			for i := c.y; i <= c2.y; i++ {
				coords = append(coords, Coordinate{x: c.x, y: i})
			}
		}
	} else if c.y == c2.y {
		if c.x > c2.x {
			for i := c2.x; i <= c.x; i++ {
				coords = append(coords, Coordinate{x: i, y: c.y})
			}
		} else {
			for i := c.x; i <= c2.x; i++ {
				coords = append(coords, Coordinate{x: i, y: c.y})
			}
		}
	}
	return coords
}

func (c Coordinate) diagonalLeft() Coordinate {
	return Coordinate{
		x: c.x - 1,
		y: c.y + 1,
	}
}

func (c Coordinate) diagonalRight() Coordinate {
	return Coordinate{
		x: c.x + 1,
		y: c.y + 1,
	}
}

func (c Coordinate) down() Coordinate {
	return Coordinate{
		x: c.x,
		y: c.y + 1,
	}
}
