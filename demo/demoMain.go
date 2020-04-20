package main

import (
	"./circle.go"
	"./controller.go"
	"./demoInterface.go"
	"./rectangle.go"
)

func main() {
	r := rectangle{length: 3, width: 2}
	measurement(r)
}
