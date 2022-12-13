package day13_test

import (
	"os"
	"testing"

	"github.com/natalieytan/aoc2022/internal/day13"
)

func Test_Part1(t *testing.T) {
	data := loadTestDataPt1(t)

	expectedResult := 13
	result := day13.Part1(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func Test_Part2(t *testing.T) {
	data := loadTestDataPt2(t)

	expectedResult := 140
	result := day13.Part2(data)

	if result != expectedResult {
		t.Errorf("expected: %d, got: %d", expectedResult, result)
	}
}

func loadTestDataPt1(t *testing.T) []day13.PacketPair {
	bytes, err := os.ReadFile("day13_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	packetPairs, err := day13.PrepareDataPt1(bytes)
	if err != nil {
		t.Fatalf(err.Error())
	}

	return packetPairs
}

func loadTestDataPt2(t *testing.T) []day13.Packet {
	bytes, err := os.ReadFile("day13_test_sample.txt")
	if err != nil {
		t.Fatalf("failure to load test data")
	}
	packetData, err := day13.PrepareData2Pt2(bytes)
	if err != nil {
		t.Fatalf(err.Error())
	}

	return packetData
}
