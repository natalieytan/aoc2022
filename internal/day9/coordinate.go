package day9

import (
	"fmt"
	"math"
)

type Coordinate struct {
	x int
	y int
}

func NewCoordinate() *Coordinate {
	return &Coordinate{}
}

func (c *Coordinate) Right() {
	c.x += 1
}

func (c *Coordinate) Left() {
	c.x -= 1
}

func (c *Coordinate) Up() {
	c.y += 1
}

func (c *Coordinate) Down() {
	c.y -= 1
}

func (c *Coordinate) MoveInDirection(d Direction) {
	switch d {
	case Up:
		c.Up()
	case Down:
		c.Down()
	case Right:
		c.Right()
	case Left:
		c.Left()
	}
}

func (c *Coordinate) MoveToTouchCoordinate(head *Coordinate) {
	if (math.Abs(float64(c.y-head.y))) <= 1 && (math.Abs(float64(c.x-head.x))) <= 1 {
		// already touching - No movement required
		return
	}

	// Move vertically if required
	if head.y > c.y {
		c.Up()
	} else if head.y < c.y {
		c.Down()
	}

	// Move horizontally if required
	if head.x > c.x {
		c.Right()
	} else if head.x < c.x {
		c.Left()
	}
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("%d-%d", c.x, c.y)
}
