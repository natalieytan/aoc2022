package day10

type OperationType int

const (
	Noop OperationType = iota
	Addx
)

type Operation struct {
	operationType OperationType
	value         int
}

func (o Operation) cyclesToComplete() int {
	switch o.operationType {
	case Noop:
		return 1
	case Addx:
		return 2
	default:
		return 0
	}
}
