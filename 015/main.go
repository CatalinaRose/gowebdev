package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", puppy)
	http.Handlefunc("/harley.jpg", dogPic)
	http.ListenAndServe(":8084", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, '<img src="toby.jpg"> ')
}

func dogPic(w http.ResponseWriter, rep *http.Request) {
	http.ServeFile(w, req, "public/img/toby.jpg")
}
