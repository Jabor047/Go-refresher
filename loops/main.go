package main

import (
	"fmt"
	"os"
)

func main(){
	println("A simple While loop")
	a := 0

	for a < len(os.Args) {
		fmt.Printf("a[%d] = %s\n", a, os.Args[a])
		a ++
	}
}