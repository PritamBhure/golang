package myutil

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message)
}

func Season(message string) {
	switch message {
	case "January", "February", "March":
		fmt.Println("The Season is Winter")
	case "April", "may", "june":
		fmt.Println("The season is summer")
	default:
		fmt.Println("Other season")
	}
}
