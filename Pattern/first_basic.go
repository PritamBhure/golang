package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {

		for Space := 4; Space >= i; Space-- {
			fmt.Print(" ")
		}

		for row := 1; row <= i; row++ {
			fmt.Print("*")

		}

		println("")
	}
}
