package main

import (
	"fmt"
	"Udemy_Go/functions_and_scopes/inner_module"
)

func add3(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Printf("My num is: %d \n", add3(4, 5, 6))
	x := add4(5, 6, 7, 8)
	fmt.Printf("My Second num is: %d\n", x)

	innermodule.Show2Nums(x, 5)

	y := "Hello"
	var g string = y + " World"
	println(g)
}
