package main

import (
	"encoding/json"
	"fmt"
)

//Cars struct
type Cars struct {
	Brand string
	Model string
	Price int
	Owner string
	Avbl  bool `json:"Availibility Status,omitempty"`
}

func main() {
	data := []Cars{
		{Brand: "BMW", Model: "GT100", Price: 5900000, Owner: "OP Singh", Avbl: true},
		{Brand: "Mercedes", Model: "S100", Price: 2200000, Owner: "DN Singh", Avbl: false},
		{Brand: "Audi", Model: "Q3", Price: 1800000, Owner: "AK Jha", Avbl: true},
		{Brand: "Hyundai", Model: "Santafe", Price: 1300000, Owner: "Harsh Madhok", Avbl: false}}

	//marshalling in specified manner
	gaadi, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n\n%s\n\n", gaadi)

	//marshalling can be read easily
	gaadi3, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n\n%s\n\n", gaadi3)

	//for unmarshalling
	var gaadi2 []struct {
		Brand string
		Model string
	}
	err = json.Unmarshal(gaadi, &gaadi2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(gaadi2, "\n\n")
}
