package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
}

func main() {
	// composite literal; struct literal
	p1 := person{"Arroyo", "Arroyo"}
	p2 := person{"GAta", "boo"}
	fmt.Println(p1)
	fmt.Println(p2)
	n := meaning()
	fmt.Println(n)
}

