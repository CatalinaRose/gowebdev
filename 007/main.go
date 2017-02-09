package main

import (
	"fmt"
)

type person struct {
	last string
	age  int
}

func main() {
	p1 := person{"Harley Quinn", 214}
	fmt.Println(p1)
	p1.foo()
}

func (p person) foo() {
	fmt.Println("Hello from", p.last)
}
