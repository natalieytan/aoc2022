package day13_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day13"
)

func Test_Part1(t *testing.T) {
	data := loadTestData(t)

	expectedResult := 13
	result := day13.Part1(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestData2(t)

	expectedResult := 140
	result := day13.Part2(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func loadTestData(t *testing.T) []day13.PacketPair {
	bytes, err := os.ReadFile("day13_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	packetPairs, err := day13.PrepareData(bytes)
	if err != nil {
		t.Fatalf(err.Error())
	}

	return packetPairs
}

func loadTestData2(t *testing.T) []day13.PacketData {
	bytes, err := os.ReadFile("day13_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	packetPairs, err := day13.PrepareData2(bytes)
	if err != nil {
		t.Fatalf(err.Error())
	}

	return packetPairs
}
