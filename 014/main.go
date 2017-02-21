package main

import (
	"fmt"
)


type person struct {
	first string
	last string
}

func main() {
	p1 := person{"Griselda", "PotheadFace"}
	p2 := person{"Coins", "drop the $$$$"}
	xp := []person{}
	xp = append(xp, p1)
	xp = append(xp, p2)
	fmt.Println(xp)
}
