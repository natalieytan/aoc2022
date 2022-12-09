package day2

func Part1(matchups []SymbolMatchup) int {
	totalScore := 0
	for _, matchup := range matchups {
		totalScore += scoreForSymbol(matchup.Player)
		totalScore += matchup.score()
	}
	return totalScore
}

func Part2(matchups []OutcomeMatchup) int {
	totalScore := 0
	for _, matchup := range matchups {
		totalScore += scoreForSymbol(matchup.requiredSymbol())
		totalScore += scoreForOutcome(matchup.Outcome)
	}
	return totalScore
}
