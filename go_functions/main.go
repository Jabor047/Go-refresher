package main

import (
	"fmt"
	"time"
)

func loopto(s string, n int) {
	for x := 0; x < n; x++ {
		fmt.Printf("%s = %d\n", s, x)
		time.Sleep(time.Millisecond)
	}
}

func main()  {
	println("Start")
	go loopto("func", 100)

	scoped := func ()  {
		for y := 0; y < 100; y++ {
			fmt.Printf("scoped = %d\n", y)
			time.Sleep(time.Millisecond)
		}
	}

	go scoped()

	go func ()  {
		for y := 0; y < 100; y++ {
			fmt.Printf("Direct = %d\n", y)
			time.Sleep(time.Millisecond)
		}
	}() // () calls the just defined function

	time.Sleep(time.Second)
	println("End")
}