package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", cold)
	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":", nil)
}

func cold(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<public src="/assets/cold.jpg">`)


}


