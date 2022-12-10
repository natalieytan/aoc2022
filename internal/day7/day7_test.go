package day7_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day7"
)

func Test_Part1(t *testing.T) {
	rootDirectory := loadRootDirectoryFromTestData(t)

	expectedResult := 95437
	result := day7.Part1(rootDirectory)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	rootDirectory := loadRootDirectoryFromTestData(t)

	expectedResult := 24933642
	result := day7.Part2(rootDirectory)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func loadRootDirectoryFromTestData(t *testing.T) *day7.Directory {
	bytes, err := os.ReadFile("day7_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	data, err := day7.PrepareData(bytes)
	if err != nil {
		t.Fatalf("failure to parse test data")
	}

	rootDirectory := day7.BuildRootDirectory(data)
	return rootDirectory
}
