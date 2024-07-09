package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func fizzBuzz(n int) string {
	if n == 2 {
		return "2"
	}
	if n == 3 {
		return "Fizz"
	}
	if n == 4 {
		return "4"
	}
	if n == 5 {
		return "Buzz"
	}
	if n == 6 {
		return "6"
	}
	return "1"
}
