package main

import (
	"fmt"
)

func threeVariable(x, y int) {
	var c int
	c = x
	x = y
	y = c

	fmt.Printf("Value Present in A is %d \n", x)

	fmt.Printf("Value Present in B is %d \n", y)
}

func withoutthreeVariable(x, y int) {
	x = x + y
	y = x - y
	x = x - y

	fmt.Printf("Value Present in A is %d \n", x)
	fmt.Printf("Value Present in B is %d \n", y)

}

func without_Aditioin_opreator(x, y int) {
	x = x ^ y
	y = x ^ y
	x = x ^ y

	fmt.Printf("Value Present in A is %d \n", x)
	fmt.Printf("Value Present in B is %d \n", y)

}

func main() {
	threeVariable(10, 20)
	withoutthreeVariable(20, 70)
	without_Aditioin_opreator(10, 40)

}
