package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

func fizzBuzz(num int) string {
	if num == 3 {
		return "Fizz"
	}
	return fmt.Sprint(num)
}
