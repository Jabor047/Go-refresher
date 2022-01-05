package main

import (
	"fmt"
)

func main() {
	sub := func(n int) int {
		return n - 7
	}

	fmt.Printf("sub(10) is %d\n", sub(10))

	a3 := AddNer(3)
	fmt.Printf("a3(6) = %d\n", a3(6))
}

func AddNer(a int) func(int) int {
	return func(b int) int{
		return a + b
	}
}