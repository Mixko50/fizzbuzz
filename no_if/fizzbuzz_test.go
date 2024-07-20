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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := FizzBuzz(tc.input)
			if got != tc.expected {
				t.Errorf("expected %s but got %s", tc.expected, got)
			}
		})
	}
}
