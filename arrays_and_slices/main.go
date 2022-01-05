package main

import "fmt"

func main(){
	a5 := [5]float64{3, 5., 2.}
	fmt.Printf("a5 = %v\n", a5)

	a10 := [10]float64{64, 56, 78}
	fmt.Printf("a10 = %v\n", a10)

	view := a5[2:]
	fmt.Printf("view = %v\n", view)

	view[0] = 5
	fmt.Printf("view (edited) = %v\n", view)
	fmt.Printf("a5 (edited) = %v\n", a5)

	view = append(view, 4)
	fmt.Printf("view (edited + 4) = %v\n", view)
	view[1] = 10
	fmt.Printf("view (edited appended) = %v\n", view)
	fmt.Printf("a5 (edited) appended = %v\n", a5)

	view = append(view, a10[:]...)
	fmt.Printf("view (edited append a10) = %v\n", view)
}