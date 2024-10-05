package main

import (
	"fmt"
	"strings"
)

func main() {

	// Remove ""
	data := "apple,orange,banana"

	parts := strings.Split(data, ",")

	fmt.Println(parts)

	//count accarangs
	str1 := "one two two three four five five five"

	count := strings.Count(str1, "two")

	println(count)

	//Trip space

	str2 := "                 How        my    name is!  "

	trim := strings.TrimSpace(str2)

	println(trim)

	//Concatination words

	str3 := "pritam"
	str4 := "Rajesh"
	str5 := "Bhure"

	fullname := strings.Join([]string{str3, str4, str5}, " ")

	println(fullname)

}
