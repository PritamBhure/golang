package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}
func main() {

	fmt.Println("Starting of the Program")

	result := add(10, 30)
	// defer keyword main function ke excution ke bad run hota hai
	// agar ekse jada defer keyword ho to woh LIFO principle follow krta hai

	defer fmt.Println("addition of two number :", result)

	defer fmt.Println("Middle of the Program")

	fmt.Println("End of the Program")
}
