package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello Programmer what is your name")
	// var name string
	// fmt.Scan(&name) //scan method spac tkhi read krta hai
	// fmt.Println("Wellcome Mr.", name

	reader := bufio.NewReader(os.Stdin)
	// BUFFER reader is reader who read data from any undefine sources like File or keyboard
	// Buffer is storage area in memory where he  store datafor temporary time
	name, _ := reader.ReadString('\n')
	fmt.Println("Wellcome Mr.", name)

}
