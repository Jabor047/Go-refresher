package main

import "fmt"

type Talker interface{
	talk() string
}

type Cat struct {}

func (c *Cat) talk() string {
	return "Meow"
}

func do_talk(c Talker) string {
	return fmt.Sprintf("The %#V says %s", c, c.talk())
}

func main(){
	println(do_talk(&Cat{}))
	println("Empty 1 -> ", empty(&Cat{}))
	println("Empty 2 ->", empty(struct{x int} {x: 5}))
}

func empty(c interface{}) string {
	// Good Practice check for abilities not type
	a, ok := c.(Talker) // check if it implements Talker

	if !ok {
		return fmt.Sprintf("The %#v cannot talk", c)
	}
	return do_talk(a)
}