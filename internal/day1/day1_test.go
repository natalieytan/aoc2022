package day1_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day1"
)

func Test_Part1(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 24000
	result := day1.Part1(data)
	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 45000
	result := day1.Part2(data)
	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func loadTestData(t *testing.T) [][]int {
	bytes, err := os.ReadFile("day1_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	data, err := day1.PrepareData(bytes)
	if err != nil {
		t.Fatalf("failure to prepare test data")
	}
	return data
}
