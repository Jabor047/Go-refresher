package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fname := "test_data/default.txt"

	if len(os.Args[1:]) > 1 {
		fname = fmt.Sprintf("test_data/%s.txt", os.Args[1])
	}

	f, err := os.OpenFile(fname, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	defer f.Close()

	f.Write([]byte("GoodBye Everybody Kife has just begun \n"))
}