package day14_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day14"
)

func Test_Part1(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 24
	result := day14.Part1(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 93
	result := day14.Part2(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func loadTestData(t *testing.T) day14.GridData {
	bytes, err := os.ReadFile("day14_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	rockGrid, err := day14.PrepareData(bytes)
	if err != nil {
		t.Fatalf(err.Error())
	}

	return rockGrid
}
