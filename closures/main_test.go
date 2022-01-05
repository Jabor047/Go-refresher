package main

import "testing"

func TestAdd(t *testing.T){
	x := AddNer(5)
	n := x(5)

	if n != 10 {
		t.Fatalf("Expected n = 10, got %d\n", n)
	}
}