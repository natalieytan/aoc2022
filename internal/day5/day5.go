package day5

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	number int
	from   int
	to     int
}

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

func FindTopOfCrateStackAfterMoving(crateStack [][]string, instructions []Instruction) string {
	moveCrateStack(crateStack, instructions)
	return topOfCrateStack(crateStack)
}

func FindTopOfCrateStackAfterMovingWithNewCrane(crateStack [][]string, instructions []Instruction) string {
	moveCrateStackWithNewCrane(crateStack, instructions)
	return topOfCrateStack(crateStack)
}

func NewCrateStack() [][]string {
	stack1 := []string{"F", "T", "C", "L", "R", "P", "G", "Q"}
	stack2 := []string{"N", "Q", "H", "W", "R", "F", "S", "J"}
	stack3 := []string{"F", "B", "H", "W", "P", "M", "Q"}
	stack4 := []string{"V", "S", "T", "D", "F"}
	stack5 := []string{"Q", "L", "D", "W", "V", "F", "Z"}
	stack6 := []string{"Z", "C", "L", "S"}
	stack7 := []string{"Z", "B", "M", "V", "D", "F"}
	stack8 := []string{"T", "J", "B"}
	stack9 := []string{"Q", "N", "B", "G", "L", "S", "P", "H"}

	return [][]string{
		stack1,
		stack2,
		stack3,
		stack4,
		stack5,
		stack6,
		stack7,
		stack8,
		stack9,
	}
}

func topOfCrateStack(crateStack [][]string) string {
	str := ""
	for _, stack := range crateStack {
		if len(stack) > 0 {
			str += stack[len(stack)-1]
		}
	}
	return str
}

func moveCrateStack(crateStack [][]string, instructions []Instruction) {
	for _, instruction := range instructions {
		move(crateStack, instruction)
	}
}

func move(stack [][]string, instruction Instruction) {
	for i := 0; i < instruction.number; i++ {
		fromStack := stack[instruction.from-1]
		if len(fromStack) > 0 {
			crate := fromStack[len(fromStack)-1]
			stack[instruction.from-1] = fromStack[:len(fromStack)-1]

			toStack := stack[instruction.to-1]
			toStack = append(toStack, crate)
			stack[instruction.to-1] = toStack
		}
	}
}

func moveCrateStackWithNewCrane(crateStack [][]string, instructions []Instruction) {
	for _, instruction := range instructions {
		moveWithNewCrane(crateStack, instruction)
	}
}

func moveWithNewCrane(stack [][]string, instruction Instruction) {
	fromStack := stack[instruction.from-1]
	cratesToMove := fromStack[len(fromStack)-(instruction.number):]
	stack[instruction.from-1] = fromStack[:len(fromStack)-(instruction.number)]

	toStack := stack[instruction.to-1]
	toStack = append(toStack, cratesToMove...)
	stack[instruction.to-1] = toStack
}
