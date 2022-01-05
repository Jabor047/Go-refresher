package main

import (
	"fmt"
	// "io"
	"log"
	"os"
)

func main() {
	
	for _, a := range(os.Args[1:]) {
		fmt.Printf("----  %s  ----\n", a)
		f, err := os.Open(a)
		if err != nil {
			log.Fatal(err.Error())
			continue
		}

		// Easy Way
		// io.Copy(os.Stdout, f)

		//Hard Way
		var buf = make([]byte, 5)

		for n, err := f.Read(buf); n > 0; n, err = f.Read(buf) {
			if err != nil {
				fmt.Printf("Err = %s\n", err.Error())
			}

			// Do what you want with buf contents
			fmt.Printf("Buf = %s\n", buf[: n])
		}
		f.Close()
	} 
}