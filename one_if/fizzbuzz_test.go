package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	type testCase struct {
		name     string
		input    int
		expected string
	}

	testCases := []testCase{
		{
			"when input is 1 should return 1",
			1,
			"1",
		},
		{
			"when input is 2 should return 2",
			2,
			"2",
		},
		{
			"when input is 3 should return 3",
			3,
			"Fizz",
		},
		{
			"when input is 4 should return 4",
			4,
			"4",
		},
		{
			"when input is 5 should return 5",
			5,
			"Buzz",
		},
		{
			"when input is 6 should return 6",
			6,
			"Fizz",
		},
		{
			"when input is 7 should return 7",
			7,
			"7",
		},
		{
			"when input is 8 should return 8",
			8,
			"8",
		},
		{
			"when input is 9 should return 9",
			9,
			"Fizz",
		},
		{
			"when input is 10 should return 10",
			10,
			"Buzz",
		},
		{
			"when input is 11 should return 11",
			11,
			"11",
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
