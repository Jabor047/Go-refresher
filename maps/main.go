package main

import "fmt"

func main(){
	emap := make(map[string] int)
	fmt.Printf("%v\n", emap)

	m := map[string]int{"k": 11}
	m["p"] = 7
	m["v"] = 11
	m["no"] = 0

	fmt.Printf("m=%v\n", m)

	for k, v := range m {
		fmt.Printf("m[%s] = %d\n", k, v)
	}

	delete(m, "no")
	fmt.Printf("m=%v\n", m)

	fmt.Printf("m[p] = %d\n", m["p"])
	
	//checking whether key exist in a map
	for _, v := range []string{"go", "no", "p"} {
		atk, ok := m[v]
		start := "Something"
		if !ok {
			start = "Nothing"
		}

		fmt.Printf("%s at '%s' : %d\n", start, v, atk)
	} 
}