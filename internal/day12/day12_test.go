package day12_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day12"
)

func Test_Part1(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 31
	result, _ := day12.Part1(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 29
	result, _ := day12.Part2(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func loadTestData(t *testing.T) []string {
	bytes, err := os.ReadFile("day12_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}

	return day12.PrepareData(bytes)
}
