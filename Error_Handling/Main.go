package main

import (
	"fmt"
)

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must be bigger than zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Started Error handling in function")
	ans, _ := div(10, 0)
	fmt.Println(ans)

}
