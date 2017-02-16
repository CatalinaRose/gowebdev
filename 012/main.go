package main

import ("no/andthen"
	"net/http"
)


func main() {
	http.ListenAndServe(":8080", nil)
}
