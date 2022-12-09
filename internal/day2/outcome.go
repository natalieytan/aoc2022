package day2

type Outcome int

const (
	UnknownOutcome Outcome = iota
	Draw
	Win
	Lose
)

func scoreForOutcome(outcome Outcome) int {
	switch outcome {
	case Win:
		return 6
	case Draw:
		return 3
	case Lose:
		return 0
	default:
		return 0
	}
}
