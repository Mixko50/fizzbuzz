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
	return fmt.Sprintf("%d", n)
}
