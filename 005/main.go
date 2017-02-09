package main

import (
	"fmt"
)

type person struct {
	fname string
	age   int
	prefs []string
}

func main() {
	p1 := person{
		"Cat",
		33,
		[]string{"I", "told", "you", "so"},
	}
	fmt.Println(p1.prefs)
}
