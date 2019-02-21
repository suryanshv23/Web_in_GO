package main

import (
	"fmt"
	"net/http"
)

func webWriter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	//Iterate over all header fields
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}
}

func main() {
	http.HandleFunc("/abcd", webWriter)
	http.ListenAndServe(":8000", nil)
}
