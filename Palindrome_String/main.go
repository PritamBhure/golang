package main

import (
	"fmt"
	"strings"
)

func Reverse_String(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	n := len(s)

	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	Result := Reverse_String("Vaibhav")
	if Result {
		fmt.Println("Given String is Palindome")
	} else {
		fmt.Println("Given String is Not Palindome")

	}

}
