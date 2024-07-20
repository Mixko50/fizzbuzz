package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func FizzBuzz(n int) string {
	output := ""
	output += map[bool]string{true: "Fizz"}[n%3 == 0]
	output += map[bool]string{true: fmt.Sprint(n)}[output == ""]
	return output
}
