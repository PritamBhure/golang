package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 42

	println(num)
	fmt.Printf("type of num is %T\n", num)

	// Data conversion int to float

	data := float64(num)
	fmt.Printf("type of data is %T\n", data)

	//int to string

	num = 123
	str := strconv.Itoa(num)
	fmt.Printf("type of str is %T\n", str)
	fmt.Printf("type of str is %s\n", str)

	// string to int

	Num_str := "12345"

	num1, _ := strconv.Atoi(Num_str)
	fmt.Printf("type of num1 is %T\n", num1)
	fmt.Printf("type of num1 is %d\n", num1)

	//string to float
	println("string to float")
	Num_string := "5.9"

	num2, _ := strconv.ParseFloat(Num_string, 64)
	fmt.Printf("type of num1 is %T\n", num2)
	fmt.Printf("type of num1 is %f\n", num2)

}
