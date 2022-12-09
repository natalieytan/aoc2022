package day5_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day5"
)

func Test_Part1(t *testing.T) {
	data := loadTestData(t)

	expectedResult := "CMZ"
	result := day5.Part1(sampleTestCrateStack(), data)

	if result != expectedResult {
		t.Errorf("expected: %s, got: %s", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestData(t)

	expectedResult := "MCD"
	result := day5.Part2(sampleTestCrateStack(), data)

	if result != expectedResult {
		t.Errorf("expected: %s, got: %s", expectedResult, result)
	}
}

func sampleTestCrateStack() day5.CrateStack {
	stack1 := []string{"Z", "N"}
	stack2 := []string{"M", "C", "D"}
	stack3 := []string{"P"}

	crates := [][]string{
		stack1,
		stack2,
		stack3,
	}

	return day5.CrateStack{
		Crates: crates,
	}
}

func loadTestData(t *testing.T) []day5.Instruction {
	bytes, err := os.ReadFile("day5_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	data, err := day5.PrepareData(bytes)
	if err != nil {
		t.Fatalf("failure to parse test data")
	}
	return data
}
