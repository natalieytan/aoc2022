package day11_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day11"
)

func Test_Part1(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 10605
	result := day11.Part1(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 2713310158
	result := day11.Part2(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func loadTestData(t *testing.T) []*day11.Monkey {
	bytes, err := os.ReadFile("day11_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	data, err := day11.PrepareData(bytes)
	if err != nil {
		t.Fatalf("failure to parse test data")

	}

	return data
}
