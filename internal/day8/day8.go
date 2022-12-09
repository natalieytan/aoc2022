package day8

func Part1(forest Forest) int {
	total := 0
	total += forest.markAndCountVisibleTreesFromLeft()
	total += forest.markAndCountVisibleTreesFromRight()
	transposedForest := forest.transpose()
	total += transposedForest.markAndCountVisibleTreesFromLeft()
	total += transposedForest.markAndCountVisibleTreesFromRight()
	return total
}

func Part2(forest Forest) int {
	maxScore := 0
	for row, treeRow := range forest.Trees {
		for col, _ := range treeRow {
			score := forest.scenicScoreForTreeAtCoordinate(row, col)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}
