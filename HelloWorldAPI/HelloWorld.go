package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//API skdjfh
type API struct {
	Message string `json:"First_Message"`
	// `...` just to give a new name to the field of struct in API's
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		message := API{"Hello, world!"}
		output, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Something went wrong!")
		}
		fmt.Fprintf(w, string(output))
	})
	http.ListenAndServe(":8080", nil)
}
