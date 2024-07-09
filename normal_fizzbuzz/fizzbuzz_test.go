package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type testCase struct {
		name     string
		input    int
		expected string
	}

	testCases := []testCase{
		{
			"1",
			1,
			"1",
		},
		{
			"2",
			2,
			"2",
		},
		{
			"3",
			3,
			"Fizz",
		},
		{
			"4",
			4,
			"4",
		},
		{
			"5",
			5,
			"Buzz",
		},
		{
			"6",
			6,
			"6",
		},
		{
			"7",
			7,
			"7",
		},
		{
			"8",
			8,
			"8",
		},
		{
			"9",
			9,
			"Fizz",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := fizzBuzz(tc.input)
			if actual != tc.expected {
				t.Errorf("fizzBuzz(%d) = %s; expected %s", tc.input, actual, tc.expected)
			}
		})
	}
}
