package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

func fizzBuzz(num int) string {
	output := ""
	output += map[bool]string{num%3 == 0: "Fizz"}[true]
	output += map[bool]string{num%5 == 0: "Buzz"}[true]
	if output == "" {
		output = fmt.Sprint(num)
	}
	return output
}
