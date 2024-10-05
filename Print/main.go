package main

import "fmt"

func main() {

	age := 24
	name := "Pritam"
	height := 5.9

	fmt.Println("ag:", age, "name:", name, "height:", height)

	//Print f Formatiion style
	fmt.Printf("age is %d\n ", age)
	fmt.Printf("age is %.2f\n ", height) // show result upto two decimal after point
	fmt.Printf("Type of age is %T\n ", age)
	fmt.Printf("Name is %s\n ", name)

	fmt.Printf("age is %d, name is %s, Height is %.3f", age, name, height)

}
