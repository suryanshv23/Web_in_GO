package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gorilla"
	"net/http"
)

type user struct {
	name   string
	place  string
	branch string
	email  string
}

var database *sql.DB

func webWriter(w http.ResponseWriter, r *http.Request) {
	newUser := user{}
	newUser.name = r.FormValue("name")
	newUser.place = r.FormValue("place")
	newUser.branch = r.FormValue("branch")
	newUser.email = r.FormValue("email")
	output, err := json.Marshal(newUser)
	fmt.Println(string(output))
	createAndOpen("users")

	sql1 := "INSERT INTO users set user_nickname='" + newUser.name + "',user_email='" + newUser.email + "',user_branch='" + newUser.branch + "',user_place='" + newUser.place + "'"
	q, err := database.Exec(sql1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
}

func createAndOpen(name string) *sql.DB {

	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/user/create", webWriter).Methods("GET")
	http.ListenAndServe(":3306", r)
}
