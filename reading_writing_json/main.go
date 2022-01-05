package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"os"
)

type Person struct {
	Name string `json:"username"`
	Age int `json:"age"`
}

func main() {

	var people []Person

	err := json.NewDecoder(os.Stdin).Decode(&people)
	if err != nil {
		log.Fatal(err.Error())
	}

	for k := range people {
		people[k].Age += 1
	}

	json.NewEncoder(os.Stdout).Encode(people)
}