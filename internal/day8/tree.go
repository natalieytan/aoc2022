package day8

type Tree struct {
	Height          int
	MarkedAsVisible bool
}

func (t *Tree) scenicScoreForView(view []*Tree) int {
	score := 0
	for _, tree := range view {
		score += 1
		if tree.Height >= t.Height {
			return score
		}
	}
	return score
}
