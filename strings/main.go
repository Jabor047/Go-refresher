package main

import (
	"fmt"
 	"strings"
)

func main()  {
	s := "你好 is hello in 中国"
	cat := s + ". I like food"

	fmt.Println(cat)
	for k, v := range s {
		fmt.Printf("[%d] = %c\n", k, v)
	}

	j := strings.Join([]string{"cat", "dog", "parrot"}, ",")
	println(j)

	k := strings.Split(j, ",")
	fmt.Printf("resplit j = %v\n", k)
}