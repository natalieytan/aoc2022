package day5

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func PrepareData(input []byte) ([]Instruction, error) {
	data := string(input)
	rawInstructions := strings.Split(data, "\n")
	instructions := make([]Instruction, len(rawInstructions))
	for i, rawInstruction := range rawInstructions {
		instruction, err := parseInstruction(rawInstruction)
		if err != nil {
			return instructions, err
		}
		instructions[i] = instruction
	}
	return instructions, nil
}

func parseInstruction(instruction string) (Instruction, error) {
	instructionExp := regexp.MustCompile(`move (?P<num>\d+) from (?P<from>\d+) to (?P<to>\d+)`)
	match := instructionExp.FindStringSubmatch(instruction)
	if len(match) != 4 {
		return Instruction{}, errors.New("error parsing instruction")
	}
	num, _ := strconv.Atoi(match[1])
	from, _ := strconv.Atoi(match[2])
	to, _ := strconv.Atoi(match[3])

	return Instruction{
		number: num,
		from:   from,
		to:     to,
	}, nil
}
