package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type postOffice struct {
	Name    string
	Taluk   string
	Region  string
	Country string
}

func main() {
	data, err := http.Get("http://postalpincode.in/api/pincode/221011")
	if err != nil {
		fmt.Printf("The http request has a error : %s", err)
	} else {
		read, _ := ioutil.ReadAll(data.Body)
		//change the type of the struct
		var po []postOffice
		err = json.Unmarshal(read, &po)
		if err != nil {
			fmt.Printf("%s", err)
		}
		for _, y := range po {
			fmt.Println(y.Name)
		}
	}

}
