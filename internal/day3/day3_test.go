package day3_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day3"
)

func Test_Part1(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 157
	result := day3.Part1(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 70
	result, err := day3.Part2(data)
	if err != nil {
		t.Errorf("weird input - %s", err.Error())
	}

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func loadTestData(t *testing.T) []string {
	bytes, err := os.ReadFile("day3_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	return day3.PrepareData(bytes)
}
