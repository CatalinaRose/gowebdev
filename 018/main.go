package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foolio)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8083", nil)
}

func foolio(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Println(req.URL.Path)
	fmt.Fprintln(w, "go fukin suck a dick!! lol")
}

