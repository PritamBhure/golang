package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// to store response in struc

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json: "completed"`
}

func performGetResponse() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Responses", res.StatusCode)
	}

	if err != nil {
		fmt.Println("Error while Getting Response", err)
		return
	}

	/*
		res_data, error := ioutil.ReadAll(res.Body)

		if error != nil {
			fmt.Println("Error while Getting Response", error)
			return
		}
		defer res.Body.Close()

		fmt.Println(string(res_data))
	*/

	var todo Todo
	//NewDecoder kay karega decode karega normal object main or variable ke address(&todo) main save krega
	error := json.NewDecoder(res.Body).Decode(&todo)
	if error != nil {
		fmt.Println("Error while Getting Response", error)
		return
	}

	fmt.Println("Decode json data", todo)

}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Workout",
		Completed: true,
	}

	// hum jab bhi koi data bhejte hai or add krte hai to use json fomate main bhejte hai

	// for that we use marshal to convert data in to json data

	jsondata, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error while Posting request", err)
		return
	}

	fmt.Println("Marshal Response", jsondata)

	println()

	//convert json data into string
	jsonstring := string(jsondata)
	fmt.Println("Convert into sring", jsonstring)

	println()
	// convert string data to reader
	jsonreader := strings.NewReader(jsonstring)
	fmt.Println("Convert sring into readable Form", jsonreader)

	myurl := "https://jsonplaceholder.typicode.com/todos"

	// send post request to the server
	res, error := http.Post(myurl, "application/json", jsonreader)
	if error != nil {
		fmt.Println("Error while  sending Post Request to the server", error)
		return
	}

	defer res.Body.Close()

	// hum main wo response mile hai use redable formate krna hoga read krne keliye

	/*	var res_data Todo
		//NewDecoder kay karega decode karega normal object main or variable ke address(&todo) main save krega
		error1 := json.NewDecoder(res.Body).Decode(&res_data)
		if error1 != nil {
			fmt.Println("Error while Decoding data", error1)
			return
		}

		fmt.Println("Decode json data", res_data)
	*/

	readble_data, error1 := ioutil.ReadAll(res.Body)

	if error1 != nil {
		println("error in reading data", error1)
	}
	println()
	fmt.Println("Response", string(readble_data))
}

func performUpdateRequset() {

	todo := Todo{
		UserID:    23056,
		Title:     " Gym Workout",
		Completed: false,
	}

	// hum jab bhi koi data bhejte hai or add krte hai to use json fomate main bhejte hai

	// for that we use marshal to convert data in to json data
	// convert json // marshall data  // Encoding data
	jsondata, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error while Marshalling data", err)
		return
	}

	fmt.Println("Encoded data into json formate ", jsondata)

	myurl := "https://jsonplaceholder.typicode.com/todos/1"

	println()
	//convert json data into string
	jsonstring := string(jsondata)
	fmt.Println("Convert into sring", jsonstring)

	println()
	// convert string data to reader
	jsonreader := strings.NewReader(jsonstring)
	fmt.Println("Convert sring into readable Form", jsonreader)

	// send post request to the server
	request, error := http.NewRequest(http.MethodPut, myurl, jsonreader)
	if error != nil {
		fmt.Println("Error while  sending Post Request to the server", error)
		return
	}
	request.Header.Set("Content-type", "application/json") // header hum fir tab banayege jab data bhejna ho

	// send request to the server
	Client := http.Client{}
	res, err := Client.Do(request)
	if err != nil {
		println("Error while sending request")
		return
	}

	defer res.Body.Close()

	readble_data, error1 := ioutil.ReadAll(res.Body)

	if error1 != nil {
		println("error in reading data", error1)
	}
	println()
	fmt.Println("Updated Response", string(readble_data))
}

func performDeletRequsent() {

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	// send post request to the server
	request, error := http.NewRequest(http.MethodDelete, myurl, nil)
	if error != nil {
		fmt.Println("Error while  sending Post Request to the server", error)
		return
	}
	request.Header.Set("Content-type", "application/json") // header hum fir tab banayege jab data bhejna ho

	// send request to the server
	Client := http.Client{}
	res, err := Client.Do(request)
	if err != nil {
		println("Error while sending request")
		return
	}

	defer res.Body.Close()

	fmt.Println("Statust code", res.StatusCode)
	if res.StatusCode == http.StatusOK {
		println("data deleted sucessfully ")
	}
}

func main() {
	print("Learning CRUD opreation in Golang")
	//	performPostRequest()
	//performUpdateRequset()
	performDeletRequsent()

}
