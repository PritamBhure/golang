package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello World")
	time.Sleep(3000 * time.Millisecond)

	fmt.Println("End hello function succsesfully")
}

func hi() {
	fmt.Println("Hi")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("Lerning Goroutine")
	go hello()
	go hi()
	time.Sleep(4000 * time.Millisecond)

}
