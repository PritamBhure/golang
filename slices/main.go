package main

//Slices is similer to array but in slice we can changes the size if slice dynamically

import (
	"fmt"
)

func main() {
	fmt.Println("Learning Slice in GOLANG")
	//number := []int{1, 2, 3, 4, 5}
	// fmt.Println("Values in number slice", number)
	// fmt.Printf("The datatypoe of data presenr is slice %T\n", number)
	// number = append(number, 10, 20, 30, 40, 50, 60, 70)
	// fmt.Println("Values after append in  number slice", number)

	// name := []string{}
	// fmt.Print(name)

	number := make([]int, 3, 5)
	number = append(number, 4)
	number = append(number, 56)
	number = append(number, 100)
	number = append(number, 4)
	number = append(number, 56)
	number = append(number, 100)
	number = append(number, 200)
	number = append(number, 300)

	fmt.Println("Slice:", number)
	fmt.Println("Length:", len(number))
	fmt.Println("Capacity:", cap(number))

	for i := 0; i < len(number)-1; i++ {
		fmt.Println("data Present in number is :", number[i])
	}

}
