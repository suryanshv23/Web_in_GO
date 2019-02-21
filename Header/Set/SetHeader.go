package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "close")
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header Field %s, Value %s \n", k, v)
		}

	})
	http.ListenAndServe(":8080", nil)
}
