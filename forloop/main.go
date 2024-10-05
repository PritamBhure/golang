package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("number is :", i)

	}

	counter := 0 //infinite loop
	for {
		fmt.Println("Counter no is :", counter)
		counter++
		if counter == 5 {
			break
		}
	}

	// range keywords returns index and value

	number := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	for index, value := range number {
		fmt.Printf("Index of Number is %d,and values is %d\n", index, value)
	}

	word := "Hello world"
	for index, value := range word {
		fmt.Printf("Index of Number is %d,and values is %c\n", index, value)
	}

}
