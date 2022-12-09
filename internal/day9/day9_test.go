package day9_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day9"
)

func Test_Part1(t *testing.T) {
	data := loadTestData1(t)

	expectedResult := 13
	result := day9.Part1(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	tests := []struct {
		inputLoader func(t *testing.T) []day9.Motion
		expected    int
	}{
		{
			inputLoader: loadTestData1,
			expected:    1,
		},
		{
			inputLoader: loadTestData2,
			expected:    36,
		},
	}

	for _, tt := range tests {
		data := tt.inputLoader(t)
		result := day9.Part2(data)
		if result != tt.expected {
			t.Errorf("expected: %d, got: %d", tt.expected, result)
		}
	}
}

func loadTestData1(t *testing.T) []day9.Motion {
	bytes, err := os.ReadFile("day9_test_sample1.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	data, err := day9.PrepareData(bytes)
	if err != nil {
		t.Fatalf("failure to parse test data")

	}

	return data
}

func loadTestData2(t *testing.T) []day9.Motion {
	bytes, err := os.ReadFile("day9_test_sample2.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	data, err := day9.PrepareData(bytes)
	if err != nil {
		t.Fatalf("failure to parse test data")

	}

	return data
}
