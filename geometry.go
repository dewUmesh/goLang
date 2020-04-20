package main

import (
	"fmt"
	"math"
)

type geometricOperation interface {
	area() float64
}

type rect struct {
	length , width float64
}

type cricle struct {
	radius float64
}

func (c cricle) area() float64{
	return math.Pi * c.radius
}

func (r rect) area() float64 {
	return r.length * r.width
}

func measurement(g geometricOperation)  {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {

	r := rect{length: 2,width: 3}
	measurement(r)
	fmt.Println("done ......")
}
