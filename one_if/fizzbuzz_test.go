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
