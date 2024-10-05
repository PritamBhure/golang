package main

import "fmt"

func main() {
	fmt.Println("Learning array in GoLANG")

	// Array Declaration then initalization

	var name [5]string
	name[4] = "Pritam"
	name[1] = "Bhure"

	fmt.Println("Name of the person is :", name)

	//Direct Inistalization

	var number = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers :", number)
	fmt.Println("Total elements in array :", len(number))              // length of array
	fmt.Println("Value in 1 index is of numnber array is ", number[1]) //print value index wise

	var price [5]string
	fmt.Println("value of an array is", price)
	fmt.Printf("value of an array is %q\n", price)

	fmt.Printf("value of name is %q\n", name)

}
