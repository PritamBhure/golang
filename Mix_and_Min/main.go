package main

import "fmt"

// Function that takes a slice as an argument
func printSlice(nums []int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func MaxElement(nums []int) {
	MaxElement := 0
	for i := 0; i < len(nums); i++ {
		if MaxElement < nums[i] {
			MaxElement = nums[i]
		}

	}
	print(MaxElement)
}

func main() {
	// Create a slice
	numbers := []int{1, 2, 3, 4, 5}

	// Pass the slice to the function
	//printSlice(numbers)
	MaxElement(numbers)
}
