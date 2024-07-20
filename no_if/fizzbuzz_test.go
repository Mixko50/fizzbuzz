package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
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
			"when input is 3 should return Fizz",
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
			"when input is 6 should return Fizz",
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
			"when input is 9 should return Fizz",
			9,
			"Fizz",
		},
		{
			"when input is 10 should return Buzz",
			10,
			"Buzz",
		},
		{
			"when input is 11 should return 11",
			11,
			"11",
		},
		{
			"when input is 12 should return Fizz",
			12,
			"Fizz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.input); got != tt.expected {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.expected)
			}
		})
	}
}
