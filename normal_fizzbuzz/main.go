package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func fizzBuzz(n int) string {
	if n == 3 {
		return "Fizz"
	}
	return fmt.Sprintf("%d", n)
}
