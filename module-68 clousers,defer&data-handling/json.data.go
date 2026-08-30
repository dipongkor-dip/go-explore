package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"` // struct tag
	Age  int    `json:"-"`    // hide age
	City string `json:"city"`
}

func json_data_func() {
	var p = person{Name: "Sunny", Age: 23, City: "Dhaka"}

	var rawJson, err = json.Marshal(p)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(string(rawJson)) // {"name":"Sunny","city":"Dhaka"}
	}

	var p2 person
	var json_text = `{"name":"Sunny","city":"Dhaka"}`

	var error = json.Unmarshal([]byte(json_text), &p2)

	if error != nil {
		fmt.Println("Error", error)
	} else {
		fmt.Printf("%+v\n", p2) // {Name:Sunny Age:0 City:Dhaka}
	}
}
