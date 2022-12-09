package day4_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day4"
)

func Test_Part1(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 2
	result := day4.Part1(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 4
	result := day4.Part2(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func loadTestData(t *testing.T) []day4.AssignmentPair {
	bytes, err := os.ReadFile("day4_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	data, err := day4.PrepareData(bytes)
	if err != nil {
		t.Fatalf("failure to parse test data")
	}
	return data
}
