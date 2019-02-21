package main

import (
	"encoding/json"
	"gorilla"
	"net/http"
	"strconv"
)

// Student details structure
type student struct {
	RollNo    int `json:"rollno"`
	Name      *Names
	Branch    string
	ContactNo int `json:"contactno"`
}

//Names of students structure
type Names struct {
	FirstName string `json:"firstname,int"`
	LastName  string `json:"lastname,int"`
}

//slice of type student to store students name
var students []student

//to get details of all students of the slice
func allStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)

}

//to get detail of specific detail of the student by giving roll no. in url
func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, y := range students {
		if strconv.Itoa(y.RollNo) == params["rollno"] {
			json.NewEncoder(w).Encode(y)
		}
	}
}

//to add a student in slice by giving all its details
func addStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var stu student
	_ = json.NewDecoder(r.Body).Decode(&stu)

	students = append(students, stu)
	json.NewEncoder(w).Encode(students)
}

//to update details of student
func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var stu student
	for x, y := range students {
		if strconv.Itoa(y.RollNo) == params["rollno"] {
			students = append(students[:x], students[x+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&stu)
			students = append(students, stu)
			break
		}
	}
	json.NewEncoder(w).Encode(students)

}

//to delete a specific student from the slice
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for x, y := range students {
		if strconv.Itoa(y.RollNo) == params["rollno"] {
			students = append(students[:x], students[x+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}

func main() {
	//adding content to the slice called students
	students = append(students, student{RollNo: 85, Branch: "IT", ContactNo: 9879585464, Name: &Names{FirstName: "Suyash", LastName: "Agrawal"}})
	students = append(students, student{RollNo: 82, Branch: "IT", ContactNo: 4546646464, Name: &Names{FirstName: "Srijan", LastName: "Agrawal"}})
	students = append(students, student{RollNo: 84, Branch: "IT", ContactNo: 8798798978, Name: &Names{FirstName: "Suryansh"}})
	students = append(students, student{RollNo: 56, Branch: "CSE", ContactNo: 2333213211, Name: &Names{FirstName: "Rahul", LastName: "Rajput"}})

	//using gorilla mux as the router
	r := mux.NewRouter()
	r.HandleFunc("/students", allStudents).Methods("GET")
	r.HandleFunc("/getStudent/{rollno}", getStudent).Methods("GET")
	r.HandleFunc("/update/{rollno}", updateStudent).Methods("PUT")
	r.HandleFunc("/delete/{rollno}", deleteStudent).Methods("DELETE")
	r.HandleFunc("/add", addStudent).Methods("POST")

	http.ListenAndServe(":8080", r)

}
