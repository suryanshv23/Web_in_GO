package main

import (
	"fmt"
	"net/http"

	"gorilla"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}/{finalAddress}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		addr := vars["finalAddress"]

		fmt.Fprintf(w, "<b>You've requested the book: %s<br>on page %s<br>final adress %s</b>", title, page, addr)
	})

	http.ListenAndServe(":8080", r)
}
