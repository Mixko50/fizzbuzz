package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

func fizzBuzz(num int) string {
	output := ""
	output += map[bool]string{num%3 == 0: "Fizz", num%3 != 0: fmt.Sprint(num)}[true]
	if num%5 == 0 {
		return "Buzz"
	}
	return output
}
