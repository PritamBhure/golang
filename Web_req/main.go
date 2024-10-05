package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while Getting Response", err)
		return
	}

	// ek bar body se hum main response aagay to hum resorces ko free kr denge
	defer res.Body.Close()

	fmt.Printf("Type of the Response we get througth api is: %T\n", res)
	//fmt.Println(res)
	data, error := ioutil.ReadAll(res.Body) // data main  array of bytes mil raha hai

	if error != nil {
		fmt.Println("Error while Reading Response", error)
		return
	}

	fmt.Println(string(data))

}
