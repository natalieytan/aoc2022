package day12

import "strings"

func PrepareData(input []byte) []string {
	data := string(input)

	return strings.Split(data, "\n")
}
