package day5

type CrateStack struct {
	Crates [][]string
}

func NewCrateStack() CrateStack {
	stack1 := []string{"F", "T", "C", "L", "R", "P", "G", "Q"}
	stack2 := []string{"N", "Q", "H", "W", "R", "F", "S", "J"}
	stack3 := []string{"F", "B", "H", "W", "P", "M", "Q"}
	stack4 := []string{"V", "S", "T", "D", "F"}
	stack5 := []string{"Q", "L", "D", "W", "V", "F", "Z"}
	stack6 := []string{"Z", "C", "L", "S"}
	stack7 := []string{"Z", "B", "M", "V", "D", "F"}
	stack8 := []string{"T", "J", "B"}
	stack9 := []string{"Q", "N", "B", "G", "L", "S", "P", "H"}

	crates := [][]string{
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

	return CrateStack{
		Crates: crates,
	}
}

func (crateStack CrateStack) top() string {
	str := ""
	for _, stack := range crateStack.Crates {
		if len(stack) > 0 {
			str += stack[len(stack)-1]
		}
	}
	return str
}

func (crateStack CrateStack) moveAll(instructions []Instruction) {
	for _, instruction := range instructions {
		crateStack.moveOne(instruction)
	}
}

func (stack CrateStack) moveOne(instruction Instruction) {
	for i := 0; i < instruction.number; i++ {
		fromStack := stack.Crates[instruction.from-1]
		if len(fromStack) > 0 {
			crate := fromStack[len(fromStack)-1]
			stack.Crates[instruction.from-1] = fromStack[:len(fromStack)-1]

			toStack := stack.Crates[instruction.to-1]
			toStack = append(toStack, crate)
			stack.Crates[instruction.to-1] = toStack
		}
	}
}

func (crateStack CrateStack) moveAllCrateMover9001(instructions []Instruction) {
	for _, instruction := range instructions {
		crateStack.moveOneCrateMover9001(instruction)
	}
}

func (stack CrateStack) moveOneCrateMover9001(instruction Instruction) {
	fromStack := stack.Crates[instruction.from-1]
	cratesToMove := fromStack[len(fromStack)-(instruction.number):]
	stack.Crates[instruction.from-1] = fromStack[:len(fromStack)-(instruction.number)]

	toStack := stack.Crates[instruction.to-1]
	toStack = append(toStack, cratesToMove...)
	stack.Crates[instruction.to-1] = toStack
}
