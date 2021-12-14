package main

import (
	"fmt"
	"log"
)

func safeDiv(a, b int) (int, error) {
	if b == 0{
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil	
}

func safeDiv2(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivZero{n: a}
	}
	return a / b, nil
}

type DivZero struct {
	n int
}

func (d *DivZero) Error() string{
	return fmt.Sprintf("Cannot divide %d by zero", d.n)
}

func (d *DivZero) Num() int {
	return d.n
}

func main(){
	r, e := safeDiv2(34, 2)
	if e != nil {
		n, ok := e.(interface{Num() int})
		if ok{
			// Check for interfaces not type
			fmt.Printf("Number Error %d\n", n.Num())
		}
		log.Fatal(fmt.Sprintf("Error:%s\n", e.Error()))
	}

	fmt.Printf("r = :%d\n", r)
}

