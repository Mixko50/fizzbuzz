package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

func fizzBuzz(num int) string {
	if num == 2 {
		return "2"
	}
	return "1"
}
