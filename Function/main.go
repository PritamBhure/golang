package main

import "fmt"

func simplefunction() {
	fmt.Println("This is my sinple Function")

}

func add(a, b int) int {
	return a + b
}

func Multiplication(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("We are learning functiona in go lang")
	simplefunction()

	ans := add(3, 5)
	fmt.Println("The sum of two number is ", ans)

	show := Multiplication(5, 5)
	fmt.Println("The sum of two number is ", show)
}
