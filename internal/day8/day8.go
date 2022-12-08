package day8

import (
	"errors"
	"strconv"
	"strings"
)

type Tree struct {
	Height          int
	MarkedAsVisible bool
}

func PrepareData(input []byte) ([][]*Tree, error) {
	data := string(input)
	rawTreeRows := strings.Split(data, "\n")
	forest := make([][]*Tree, len(rawTreeRows))

	for i, rawTreeRow := range rawTreeRows {
		treeRow, err := parseTreeRow(rawTreeRow)
		if err != nil {
			return forest, err
		}
		forest[i] = treeRow
	}
	return forest, nil
}

func parseTreeRow(rawTreeRow string) ([]*Tree, error) {
	rawTreesInRow := strings.Split(rawTreeRow, "")
	treesInRow := make([]*Tree, len(rawTreesInRow))
	for i, rawTree := range rawTreesInRow {
		treeHeight, err := strconv.Atoi(rawTree)
		if err != nil {
			return treesInRow, errors.New("error parsing tree row")
		}
		treesInRow[i] = &Tree{
			Height:          treeHeight,
			MarkedAsVisible: false,
		}
	}
	return treesInRow, nil
}

func VisibleTreesOutsideGrid(forest [][]*Tree) int {
	total := 0

	total += visibleFromLeft(forest)
	total += visibleFromRight(forest)
	transposedForest := transposeForest(forest)
	total += visibleFromLeft(transposedForest)
	total += visibleFromRight(transposedForest)
	return total
}

func visibleFromLeft(forest [][]*Tree) int {
	total := 0
	for _, treeRow := range forest {
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

func visibleFromRight(forest [][]*Tree) int {
	total := 0
	for _, treeRow := range forest {
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

func transposeForest(forest [][]*Tree) [][]*Tree {
	transposedForest := make([][]*Tree, len(forest[0]))

	for _, treeRow := range forest {
		for i, tree := range treeRow {
			transposedForest[i] = append(transposedForest[i], tree)
		}
	}

	return transposedForest
}

func scenicScoreForTree(forest [][]*Tree, row int, col int) int {
	tree := forest[row][col]
	rowLength := len(forest[0])

	viewToRight := []*Tree{}
	if col < rowLength-1 {
		viewToRight = forest[row][col+1:]
	}
	rightScore := scenicScoreForView(tree, viewToRight)

	viewToLeft := []*Tree{}
	if col > 0 {
		left := forest[row][:col]
		for i := len(left) - 1; i >= 0; i-- {
			viewToLeft = append(viewToLeft, left[i])
		}
	}
	leftScore := scenicScoreForView(tree, viewToLeft)

	transposedForest := transposeForest(forest)
	colLength := len(transposedForest[0])

	viewToBot := []*Tree{}
	if row < colLength-1 {
		viewToBot = transposedForest[col][row+1:]
	}
	botScore := scenicScoreForView(tree, viewToBot)

	viewToTop := []*Tree{}
	if row > 0 {
		top := transposedForest[col][:row]
		for i := len(top) - 1; i >= 0; i-- {
			viewToTop = append(viewToTop, top[i])
		}
	}
	topScore := scenicScoreForView(tree, viewToTop)

	total := rightScore * leftScore * botScore * topScore
	return total
}

func scenicScoreForView(baseTree *Tree, view []*Tree) int {
	score := 0
	for _, tree := range view {
		score += 1
		if tree.Height >= baseTree.Height {
			return score
		}
	}
	return score
}

func MaxScenicScore(forest [][]*Tree) int {
	maxScore := 0

	for row, treeRow := range forest {
		for col, _ := range treeRow {
			score := scenicScoreForTree(forest, row, col)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}
