package day2

type Matchup struct {
	Opponent Symbol
	Col2     string
}

type SymbolMatchup struct {
	Opponent Symbol
	Player   Symbol
}

func (sm SymbolMatchup) score() int {
	if sm.Player == sm.Opponent {
		return 3
	} else if sm.isWinning() {
		return 6
	} else {
		return 0
	}
}
func (sm SymbolMatchup) isWinning() bool {
	winningCombos := map[Symbol]Symbol{
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
	}
	return winningCombos[sm.Player] == sm.Opponent
}

type OutcomeMatchup struct {
	Opponent Symbol
	Outcome  Outcome
}

func (om OutcomeMatchup) requiredSymbol() Symbol {
	switch om.Outcome {
	case Draw:
		return om.Opponent
	case Win:
		combos := map[Symbol]Symbol{
			Rock:     Paper,
			Paper:    Scissors,
			Scissors: Rock,
		}
		return combos[om.Opponent]
	case Lose:
		combos := map[Symbol]Symbol{
			Rock:     Scissors,
			Paper:    Rock,
			Scissors: Paper,
		}
		return combos[om.Opponent]
	default:
		return UnknownSymbol
	}
}
