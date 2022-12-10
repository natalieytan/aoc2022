package day10_test

import (
	"os"
	"strings"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day10"
)

func Test_Part1(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 13140
	result := day10.Part1(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestData(t)

	expectedRows := loadCrtRows(t)
	ledResult := day10.Part2(data)

	for i, rowResult := range ledResult {
		if rowResult != expectedRows[i] {
			t.Errorf("expected: %s got: %s", expectedRows[i], rowResult)
		}
	}

}

func loadCrtRows(t *testing.T) []string {
	bytes, err := os.ReadFile("day10_test_led.txt")
	if err != nil {
		t.Fatalf("failure to load test led screen data")
	}
	data := string(bytes)
	return strings.Split(data, "\n")
}

func loadTestData(t *testing.T) []day10.Operation {
	bytes, err := os.ReadFile("day10_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	data, err := day10.PrepareData(bytes)
	if err != nil {
		t.Fatalf("failure to parse test data")

	}

	return data
}
