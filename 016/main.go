package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", cold)
	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8086", nil)
}

func cold(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/assets/cold.jpg">`)


}


