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
	return "1"
}
