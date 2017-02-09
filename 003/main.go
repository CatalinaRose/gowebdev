package main

import (
	"fmt"
)

func main() {
	s := []int{3, 6, 69}
	for i, _ := range s {
		fmt.Println(i, "-", s[i])
	}
}
