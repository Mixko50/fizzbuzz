package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func fizzBuzz(n int) string {
	output := ""
	output += map[bool]string{true: "Fizz"}[n%3 == 0]
	output += map[bool]string{true: "Buzz"}[n%5 == 0]
	output += map[bool]string{true: fmt.Sprint(n)}[output == ""]

	return fmt.Sprint(output)
}
