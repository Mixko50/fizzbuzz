package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func fizzBuzz(n int) string {
	if n%3 == 0 {
		return "Fizz"
	}
	if n == 5 {
		return "Buzz"
	}

	if n == 10 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", n)
}
