package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   // yaha per struct ke key(IsAdult) ko key pakade ga
}

func main() {
	fmt.Println("Learning JSON in Golang")
	personvar := Person{
		Name:    "Pritam",
		Age:     24,
		IsAdult: true}
	fmt.Println("Data present in personvar:", personvar)

	//personvar into JSON ENcoding Marshalling

	jsondata, err := json.Marshal(personvar)

	if err != nil {
		fmt.Println("Error while Converting data into JSON", err)
		return
	}
	println("")
	println("Data after Encode(MArshall)")
	fmt.Println(string(jsondata))

	//Decoding Json data Unmarshalling

	var persondata Person //marshall data , address where we have to decode
	error := json.Unmarshal(jsondata, &persondata)
	if error != nil {
		fmt.Println("Error while Decoding json data", error)
		return
	}
	println("")
	fmt.Println("Person data is", persondata)
}
