package main

import "fmt"

func main() {

	num := 134
	ans := 0
	Digit := 0

	for i := 0; num > 0; i++ {
		Digit = num % 10
		ans = (ans * 10) + Digit
		num = num / 10

	}

	fmt.Println("Revers of num is:=", ans)

}
