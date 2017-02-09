package main


import (
	"fmt"
)

type person struct {
	last string
	age  int
}

func main() {
	p1 := person{"Joker", 13}
	fmt.Println(p1)
	p1.foo()
}

func (p person) foo() {
	fmt.Println("Hello from Joker, baby")
}

// func (receiver) identifier(parameters) (returns) {}

