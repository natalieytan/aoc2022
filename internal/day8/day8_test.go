package day8_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day8"
)

func Test_Part1(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 21
	result := day8.Part1(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 8
	result := day8.Part2(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func loadTestData(t *testing.T) day8.Forest {
	bytes, err := os.ReadFile("day8_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	data, err := day8.PrepareData(bytes)
	if err != nil {
		t.Fatalf("failure to parse test data")

	}

	return data
}
