package main

import "fmt"

func modifyvalueByReference(num *int) {
	*num = *num * 2
}

func main() {
	fmt.Println("Lerning Pointer in GOLANG")

	// var num int = 2

	// var ptr *int
	// ptr = &num

	num := 2

	ptr := &num

	println("Address of the container which his pointer hold ", ptr)
	println("Data present in the following container address ", *ptr)

	value := 10
	modifyvalueByReference(&value)

	println("Value  contain is :", value)

}
