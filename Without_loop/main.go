package main

import "fmt"

func show(num int) {
	if num < 1 {
		return
	}
	fmt.Printf("Number := %d\n", num)
	num--
	show(num)

}

func main() {
	show(5)
}
