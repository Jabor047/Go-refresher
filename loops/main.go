package main

import (
	"fmt"
	"os"
)

func main(){
	println("A simple While loop")
	a := 0

	for a < len(os.Args) {
		fmt.Printf("	a[%d] = %s\n", a, os.Args[a])
		a++
	}

	println("A 3 part loop")
	for a := 0; a < len(os.Args); a++ {
		fmt.Printf("	a[%d] = %s\n", a, os.Args[a])
	}

	println("A Range Loop")
	for k, v := range os.Args[1:] {
		fmt.Printf("	a[%d] = %s\n", k, v)
	}

	printArgs()
}

func printArgs(){
	println("No Condition Loop")
	a := 0;

	for {
		if a >= len(os.Args){
			fmt.Println("	a[%d] = %s\n", a, os.Args[a])
			a ++
		}
	}
}