package main

import (
	"net/http"
	"io"
	"fmt"
)


type person int

func(p person)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello FReaks, Freaks do it better")
}

func main() {
	var weed person
	http.Handle("/", weed)
	err:= http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
