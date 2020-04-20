package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	a := 10
	b := 20
	c := a + b
	fmt.Println(c)
	fmt.Println(os.Getgid())
	fmt.Println(os.Hostname())
	fmt.Println(time.Now().Clock())

	print()
}

func print(){
	fmt.Print("nside the print function")
}