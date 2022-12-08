package day2

import (
	"errors"
	"strings"
)

type Symbol int

const (
	UnknownSymbol Symbol = iota
	Rock
	Paper
	Scissors
)

type Outcome int

const (
	UnknownOutcome Outcome = iota
	Draw
	Win
	Lose
)

type SymbolMatchup struct {
	Opponent Symbol
	Player   Symbol
}

type OutcomeMatchup struct {
	Opponent Symbol
	Outcome  Outcome
}

func PrepareDataPt1(input []byte) ([]SymbolMatchup, error) {
	data := string(input)
	rawPairs := strings.Split(data, "\n")
	allSymbolMatchups := make([]SymbolMatchup, len(rawPairs))

	for i, rawPair := range rawPairs {
		charPair := strings.Split(rawPair, " ")
		if len(charPair) != 2 {
			return allSymbolMatchups, errors.New("unexpected input - more than 2 chars")
		}
		opponentSymbol, err := symbolForChar(charPair[0])
		if err != nil {
			return allSymbolMatchups, err
		}
		playerSymbol, err := symbolForChar(charPair[1])
		if err != nil {
			return allSymbolMatchups, err
		}

		allSymbolMatchups[i] = SymbolMatchup{
			Opponent: opponentSymbol,
			Player:   playerSymbol,
		}
	}

	return allSymbolMatchups, nil
}

func TotalScoreForPlayerPt1(matchups []SymbolMatchup) int {
	totalScore := 0
	for _, matchup := range matchups {
		totalScore += scoreForSymbol(matchup.Player)
		totalScore += scoreForSymbolMatchup(matchup.Player, matchup.Opponent)
	}
	return totalScore
}

func PrepareDataPt2(input []byte) ([]OutcomeMatchup, error) {
	data := string(input)
	rawPairs := strings.Split(data, "\n")
	outcomeMatchups := make([]OutcomeMatchup, len(rawPairs))

	for i, rawPair := range rawPairs {
		charPair := strings.Split(rawPair, " ")
		if len(charPair) != 2 {
			return outcomeMatchups, errors.New("unexpected more than 2")
		}
		opponentSymbol, err := symbolForChar(charPair[0])
		if err != nil {
			return outcomeMatchups, err
		}

		outcome, err := outcomeForChar(charPair[1])
		if err != nil {
			return outcomeMatchups, err
		}

		outcomeMatchups[i] = OutcomeMatchup{
			Opponent: opponentSymbol,
			Outcome:  outcome,
		}
	}

	return outcomeMatchups, nil
}

func TotalScoreForPlayerPt2(matchups []OutcomeMatchup) int {
	totalScore := 0
	for _, matchup := range matchups {
		totalScore += scoreForSymbol(symbolForOutcome(matchup.Outcome, matchup.Opponent))
		totalScore += scoreForOutcome(matchup.Outcome)
	}
	return totalScore
}

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

func scoreForSymbolMatchup(playerSymbol Symbol, opponentSymbol Symbol) int {
	if playerSymbol == opponentSymbol {
		return 3
	} else if isWinningSymbolMatchup(playerSymbol, opponentSymbol) {
		return 6
	} else {
		return 0
	}
}

func symbolForOutcome(outcome Outcome, opponentSymbol Symbol) Symbol {
	switch outcome {
	case Draw:
		return opponentSymbol
	case Win:
		combos := map[Symbol]Symbol{
			Rock:     Paper,
			Paper:    Scissors,
			Scissors: Rock,
		}
		return combos[opponentSymbol]
	case Lose:
		combos := map[Symbol]Symbol{
			Rock:     Scissors,
			Paper:    Rock,
			Scissors: Paper,
		}
		return combos[opponentSymbol]
	default:
		return UnknownSymbol
	}
}

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

func isWinningSymbolMatchup(playerSymbol Symbol, opponentSymbol Symbol) bool {
	winningCombos := map[Symbol]Symbol{
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
	}
	return winningCombos[playerSymbol] == opponentSymbol
}

func symbolForChar(char string) (Symbol, error) {
	switch char {
	case "A", "X":
		return Rock, nil
	case "B", "Y":
		return Paper, nil
	case "C", "Z":
		return Scissors, nil
	default:
		return UnknownSymbol, errors.New("unknown char - unable to map to symbol")
	}
}

func outcomeForChar(char string) (Outcome, error) {
	switch char {
	case "X":
		return Lose, nil
	case "Y":
		return Draw, nil
	case "Z":
		return Win, nil
	default:
		return UnknownOutcome, errors.New("unknown char - unable to map to outcome")
	}
}
