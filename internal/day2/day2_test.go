package day2_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day2"
)

func Test_Part1(t *testing.T) {
	data := loadTestData(t)
	matchups, err := day2.CovertToSymbolMatchups(data)
	if err != nil {
		t.Fatalf("error parsing loaded data")
	}

	expectedResult := 15
	result := day2.Part1(matchups)
	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestData(t)
	matchups, err := day2.CovertToOutComeMatchups(data)
	if err != nil {
		t.Fatalf("error parsing loaded data")
	}

	expectedResult := 12
	result := day2.Part2(matchups)
	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func loadTestData(t *testing.T) []day2.Matchup {
	bytes, err := os.ReadFile("day2_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	data, err := day2.PrepareData(bytes)
	if err != nil {
		t.Fatalf("failure to prepare test data")
	}
	return data
}
