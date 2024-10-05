package main

import (
	"Mylearnings/myutil"
	"fmt"
)

func main() {
	var day int
	fmt.Println("Enter day ")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("The day is Monday")
	case 2:
		fmt.Println("The day is Tuseday")
	case 3:
		fmt.Println("The day is Wednesday")
	case 4:
		fmt.Println("The day is Thusday")
	case 5:
		fmt.Println("The day is Friday")
	case 6:
		fmt.Println("The day is Saturday")
	default:
		fmt.Println("The day is Sunday")
	}

	myutil.Season("April")

}
