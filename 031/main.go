package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/yo", boogyman)
	http.ListenAndServe(":8028", nil)
}

func boogymanfidfidfidfwafh
