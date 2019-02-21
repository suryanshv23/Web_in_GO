package main

import (
	"fmt"
	"gorilla"
	"net/http"
)

func webWriter1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!!!")
}

func webWriter2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Myself Suryansh")
}

func webWriter3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NIT Raipur")
}

func webWriter4(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Varanasi")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", webWriter1)
	r.HandleFunc("/name", webWriter2)
	r.HandleFunc("/college", webWriter3)
	r.HandleFunc("/place", webWriter4)
	http.ListenAndServe(":8080", r)
}
