package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func fizzBuzz(n int) string {
	output := ""
	output += map[bool]string{true: "Fizz", false: fmt.Sprint(n)}[n%3 == 0]
	return fmt.Sprint(output)
}
