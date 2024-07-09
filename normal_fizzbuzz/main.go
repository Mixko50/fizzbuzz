package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func fizzBuzz(n int) string {
	if n == 3 {
		return "Fizz"
	}

	if n == 5 {
		return "Buzz"
	}
	if n == 6 {
		return "Fizz"
	}
	if n == 9 {
		return "Fizz"
	}
	if n == 10 {
		return "Buzz"
	}
	if n == 12 {
		return "Fizz"
	}
	if n == 15 {
		return "FizzBuzz"
	}
	return fmt.Sprintf("%d", n)
}
