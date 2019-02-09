package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(Multiply(20, 30))
}

// AnotherMethod : returns the sum of two parameters
func AnotherMethod(firstnum int, secondnum int) int {
	return firstnum + secondnum
}

func Subtract(firstnum int, secondnum int) int {
	return firstnum - secondnum
}