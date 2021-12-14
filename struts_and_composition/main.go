package main

import "fmt"

type Point struct {
	x int
	y int
}

type Car struct {
	Colour string
	Point // anonymous type
}

// *point is a pointer - address to where the object is in memory
func (p *Point) move(x, y int) {
	p.x += x
	p.y += y
}

func (p *Point) display() string {
	return fmt.Sprintf("{x:%d, y:%d}", p.x, p.y)
}

func (c *Car) display() string {
	return fmt.Sprintf("{Colour: %s, Location: %s}", c.Colour, c.Point.display())
}

func main(){
	p:= Point{x:5, y:8}
	p.move(3, 3)
	println(p.display())

	car := Car{Colour: "Blue", Point: Point{x:5, y:4}}
	car.move(3, 4)
	println(car.display())

}