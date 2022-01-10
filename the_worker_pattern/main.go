package main

import (
	"fmt"
	"time"
)

type SResult struct {
	s string
	done bool
}

func worker(id int, in <-chan func() string, out chan<- SResult){
	for {
		j, ok := <-in
		if !ok {
			out <- SResult{s:"", done: true}
			return
		}
		s := j()
		fmt.Printf("Worker %d : %s\n", id, s)
		out <- SResult{s:s, done: false}
	}
}

func job1(x int) func() string{
	return func() string {
		time.Sleep(time.Millisecond)
		return fmt.Sprintf("Job 1, n = %d", x)
	}
}

func job2(x int) func() string{
	return func() string {
		time.Sleep(time.Millisecond)
		return fmt.Sprintf("Job 2, n = %d", x)
	}
}

func main(){
	chin := make(chan func() string, 5)
	chres := make(chan SResult, 5)

	for x := 0; x < 5; x++ {
		go worker(x, chin, chres)
	}

	go func ()  {
		for x := 0; x < 50; x++ {
			chin <- job1(x)
			chin <- job2(x)
		}
		
		close(chin)
	}()

	doneCount := 0
	for sr := range chres {
		if sr.done {
			doneCount ++
			if doneCount >= 5 {
				break
			}
			continue
		}
		fmt.Printf("complete: s = %s\n", sr.s)
	}
	println("Done")
}