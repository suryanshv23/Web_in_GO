package main

import (
	"fmt"
	"net/http"
)

func webWriter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/abcd", webWriter)
	http.ListenAndServe(":8000", nil)
}
