package day8

import (
	"errors"
	"strconv"
	"strings"
)

func PrepareData(input []byte) (Forest, error) {
	data := string(input)
	rawTreeRows := strings.Split(data, "\n")
	trees := make([][]*Tree, len(rawTreeRows))

	for i, rawTreeRow := range rawTreeRows {
		treeRow, err := parseTreeRow(rawTreeRow)
		if err != nil {
			return Forest{}, err
		}
		trees[i] = treeRow
	}
	return Forest{
		Trees: trees,
	}, nil
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
