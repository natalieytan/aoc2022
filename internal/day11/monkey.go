package day11

type Monkey struct {
	number      int
	items       []int
	decision    MonkeyDecision
	operation   Operation
	inspections int
}

type MonkeyDecision struct {
	divisbleBy            int
	throwToIfDivisible    int
	throwToIfNotDivisible int
}

type Operation struct {
	operationType OperationType
	value1IsOld   bool
	value1        int
	value2IsOld   bool
	value2        int
}

type OperationType int

const (
	UnknownOperation OperationType = iota
	Add
	Subtract
	Multiply
	Divide
)

func (m *Monkey) itemToBeThrownToMonkeyNo(item int) int {
	if item%m.decision.divisbleBy == 0 {
		return m.decision.throwToIfDivisible
	} else {
		return m.decision.throwToIfNotDivisible
	}
}

func (m *Monkey) worryLevelAfterInspectingItem(item int) int {
	var value1 int
	if m.operation.value1IsOld {
		value1 = item
	} else {
		value1 = m.operation.value1
	}
	var value2 int
	if m.operation.value2IsOld {
		value2 = item
	} else {
		value2 = m.operation.value2
	}

	switch m.operation.operationType {
	case Add:
		return value1 + value2
	case Subtract:
		return value1 - value2
	case Multiply:
		return value1 * value2
	case Divide:
		return value1 / value2
	}

	return 0
}
