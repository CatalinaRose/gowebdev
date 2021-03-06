package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", TechNine)
	http.HandleFunc("/technine.jpg", TechNinePic)
	http.ListenAndServe(":", nil)
}

func TechNine(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<public src="technine.jpg"> `)
}

func TechNinePic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "public/technine.jpg")
}
