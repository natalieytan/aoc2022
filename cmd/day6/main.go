package main

import (
	"fmt"

	"github.com/natalieytan/aoc2022/internal/day6"
)

func main() {
	part1 := day6.CharsBeforeStartOfPacket(day6.InputStr)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := day6.CharsBeforeStartOfMessage(day6.InputStr)
	fmt.Printf("Part 1: %d\n", part2)
}
