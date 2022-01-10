package main

import (
	"fmt"
	"sync"
	"time"
)

type Quiter struct {
	n int
	sync.Mutex
}

func worker(q *Quiter, n *int, change int) {
	for x := 0; x < 100; x++ {
		q.Lock()
		*n += change
		q.Unlock()
	}

	q.Lock()
	q.n -= 1
	fmt.Printf("Worker done: q.n = %d\n", q.n)
	q.Unlock()
}

func main() {
	println("Begin")
	q := Quiter{
		n: 5,
	}
	val := 0

	for x := 0; x < 5; x ++ {
		go worker(&q, &val, x)
	}

	for {
		time.Sleep(time.Millisecond)
		q.Lock()
		if q.n == 0 {
			q.Unlock()
			break
		}
		q.Unlock()
	}

	fmt.Printf("Done Val := %d\n", val)
}