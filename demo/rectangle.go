package main

type rectangle struct {
	length, width float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}
