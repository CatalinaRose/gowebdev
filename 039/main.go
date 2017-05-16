package main

import (
	"fmt"
	"encoding/json"
)


type person struct {
	First string
	Last  string
	Items []string
}

func main() {
	p1 := person{
		First: "Gata",
		Last:  "Juegos",
		Items: []string{
			"Momma",
			"kids",
			"The Arroyo's",
		},
	}
	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1)
	fmt.Println("------------")
	fmt.Println(string(bs))
}
