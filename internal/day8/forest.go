package day8

type Forest struct {
	Trees [][]*Tree
}

func (forest Forest) markAndCountVisibleTreesFromLeft() int {
	total := 0
	for _, treeRow := range forest.Trees {
		highestinRow := -1
		for _, tree := range treeRow {
			if tree.Height > highestinRow {
				highestinRow = tree.Height
				if !tree.MarkedAsVisible {
					tree.MarkedAsVisible = true
					total += 1
				}
			}
		}
	}
	return total
}

func (forest Forest) markAndCountVisibleTreesFromRight() int {
	total := 0
	for _, treeRow := range forest.Trees {
		highestinRow := -1
		for i := len(treeRow) - 1; i >= 0; i-- {
			tree := treeRow[i]
			if tree.Height > highestinRow {
				highestinRow = tree.Height
				if !tree.MarkedAsVisible {
					tree.MarkedAsVisible = true
					total += 1
				}
			}
		}
	}
	return total
}

func (forest Forest) transpose() Forest {
	transposedTrees := make([][]*Tree, len(forest.Trees[0]))

	for _, treeRow := range forest.Trees {
		for i, tree := range treeRow {
			transposedTrees[i] = append(transposedTrees[i], tree)
		}
	}

	return Forest{
		Trees: transposedTrees,
	}
}

func (forest Forest) scenicScoreForTreeAtCoordinate(row int, col int) int {
	tree := forest.Trees[row][col]
	rightScore := tree.scenicScoreForView(forest.viewToRightFromCoordinate(row, col))
	leftScore := tree.scenicScoreForView(forest.viewToLeftFromCoordinate(row, col))
	transposedForest := forest.transpose()
	botScore := tree.scenicScoreForView(transposedForest.viewToLeftFromCoordinate(col, row))
	topScore := tree.scenicScoreForView(transposedForest.viewToRightFromCoordinate(col, row))
	total := rightScore * leftScore * botScore * topScore
	return total
}

func (forest Forest) viewToRightFromCoordinate(row int, col int) []*Tree {
	rowLength := len(forest.Trees[0])
	viewToRight := []*Tree{}
	if col < rowLength-1 {
		viewToRight = forest.Trees[row][col+1:]
	}
	return viewToRight
}

func (forest Forest) viewToLeftFromCoordinate(row int, col int) []*Tree {
	viewToLeft := []*Tree{}
	if col > 0 {
		left := forest.Trees[row][:col]
		for i := len(left) - 1; i >= 0; i-- {
			viewToLeft = append(viewToLeft, left[i])
		}
	}
	return viewToLeft
}
