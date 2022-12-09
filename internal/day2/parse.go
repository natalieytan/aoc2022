package day2

import (
	"errors"
	"strings"
)

func PrepareData(input []byte) ([]Matchup, error) {
	data := string(input)
	rawPairs := strings.Split(data, "\n")
	allMatchups := make([]Matchup, len(rawPairs))

	for i, rawPair := range rawPairs {
		charPair := strings.Split(rawPair, " ")
		if len(charPair) != 2 {
			return allMatchups, errors.New("unexpected input - more than 2 chars")
		}
		opponentSymbol, err := parseSymbolABC(charPair[0])
		if err != nil {
			return allMatchups, err
		}

		allMatchups[i] = Matchup{
			Opponent: opponentSymbol,
			Col2:     charPair[1],
		}
	}

	return allMatchups, nil
}

func CovertToSymbolMatchups(matchups []Matchup) ([]SymbolMatchup, error) {
	allSymbolMatchups := make([]SymbolMatchup, len(matchups))
	for i, matchup := range matchups {
		playerSymbol, err := parseSymbolXYZ(matchup.Col2)
		if err != nil {
			return allSymbolMatchups, err
		}
		allSymbolMatchups[i] = SymbolMatchup{
			Opponent: matchup.Opponent,
			Player:   playerSymbol,
		}
	}

	return allSymbolMatchups, nil
}

func CovertToOutComeMatchups(matchups []Matchup) ([]OutcomeMatchup, error) {
	allOutcomeMatchups := make([]OutcomeMatchup, len(matchups))
	for i, matchup := range matchups {
		outcome, err := parseOutcome(matchup.Col2)
		if err != nil {
			return allOutcomeMatchups, err
		}
		allOutcomeMatchups[i] = OutcomeMatchup{
			Opponent: matchup.Opponent,
			Outcome:  outcome,
		}
	}

	return allOutcomeMatchups, nil
}

func parseSymbolABC(char string) (Symbol, error) {
	switch char {
	case "A":
		return Rock, nil
	case "B":
		return Paper, nil
	case "C":
		return Scissors, nil
	default:
		return UnknownSymbol, errors.New("unknown char - unable to map to symbol")
	}
}

func parseSymbolXYZ(char string) (Symbol, error) {
	switch char {
	case "X":
		return Rock, nil
	case "Y":
		return Paper, nil
	case "Z":
		return Scissors, nil
	default:
		return UnknownSymbol, errors.New("unknown char - unable to map to symbol")
	}
}

func parseOutcome(char string) (Outcome, error) {
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
