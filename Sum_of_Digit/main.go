package main

import "fmt"

func sumofdigit(num int) int {
	var sum = 0
	for i := 0; num != 0; i++ {

		digit := num % 10
		sum += digit
		num = num / 10
	}
	return sum

}
func main() {
	Result := sumofdigit(156)
	fmt.Print(Result)

}
