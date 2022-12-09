package day6_test

import (
	"testing"

	"github.com/natalieytan/aoc2022/internal/day6"
)

func Test_Part1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 5,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 6,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 10,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 11,
		},
	}

	for _, tt := range tests {
		result := day6.Part1(tt.input)
		if result != tt.expected {
			t.Errorf("expected: %d, got: %d", tt.expected, result)
		}
	}
}

func Test_Part2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 19,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 23,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 23,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 29,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 26,
		},
	}

	for _, tt := range tests {
		result := day6.Part2(tt.input)
		if result != tt.expected {
			t.Errorf("expected: %d, got: %d", tt.expected, result)
		}
	}
}
