package main // this file belong to this package

import (
	"fmt"
)

func main() {
	fmt.Println("Lear GO Language Hello World")

	variable()

}

func variable() {
	var name string = "Pritam"
	var version = 56
	fmt.Println(name)
	fmt.Println(version)

	var radius float64 = 5.5
	fmt.Println(radius)

	var relationship bool = false

	fmt.Println("I am in relationship is Yes/No ", relationship)

	const pi = 3.14

	fmt.Println("The value of pi is ", pi)

	//New Type

	person := 123

	fmt.Println("New method", person)

	//to acces varibale from allover then you have to declear his name in inital letter in capital
	var Public_Variable = "The data is public"
	//This is  only visiable of this package
	var private_variable = "The data is immportend and private"

	fmt.Println(Public_Variable)
	fmt.Println(private_variable)

}
