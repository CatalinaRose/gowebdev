package main

import (
	"fmt"
)

func main() {
	s := map[string]int{"WEed": 66, "Gata": 25, "pot": 420}
	fmt.Println(s)
	x := map[string]int{"Griselda": 33, "Alexis": 10, "Emma": 8}
	for k, _ := range x {
		fmt.Println(k, "-", x)
	}
}
