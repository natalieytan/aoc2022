package day2

type Symbol int

const (
	UnknownSymbol Symbol = iota
	Rock
	Paper
	Scissors
)

func scoreForSymbol(symbol Symbol) int {
	switch symbol {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		return 0
	}
}
