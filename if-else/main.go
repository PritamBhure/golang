package main

import "fmt"

func largernum(a, b, c int) int {
	if a > b {
		fmt.Printf("%d  is greater than %d", a, b)
		return a

	} else if c > b {
		fmt.Printf("%d  is greater than %d", c, b)
		return c
	} else {
		fmt.Printf("%d  is greater than %d\n&%d\n", b, a, c)
		return b
	}

}

func main() {
	x := 4
	if x > 5 && (x >= 4 || x <= 10) {
		fmt.Println("X is Greater than five")

	} else {
		fmt.Println("X is lesser then five")

	}

	show := largernum(10, 200, 30)
	fmt.Println("The  greater among A B & c is ", show)

}
