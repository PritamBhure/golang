package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//Create New file
	// file, err := os.Create("Example.txt")
	// if err != nil {
	// 	fmt.Println("Error while Creating file", err)
	// 	return
	// }

	// content := "Hello World"
	// byte, errors := io.WriteString(file, content+"\n")
	// fmt.Println("Size of the data present in the file", byte)
	// if errors != nil {
	// 	fmt.Println("Error while Creating file", errors)
	// 	return
	// }
	// fmt.Println("Successfully created the file")

	// defer file.Close() // free the os resources end of the program

	/*
		//Open file
		File, err := os.Open("Example.txt")
		if err != nil {
			fmt.Println("Error while Opening the  file")
			return
		}
		defer File.Close()
		// buffer is a temprory storage area
		buffer := make([]byte, 1024)
		// read content from the buffer memory
		for {
			n, error := File.Read(buffer)
			if error == io.EOF {
				break
			}

			if error != nil {
				fmt.Println("Error while reading Data of the file", error)
				return
			}

			fmt.Println(string(buffer[:n]))
		}

	*/

	//Read File Data using ioutil package
	// kam kaise krta hai
	// pure File ko  ek memory main save krta hai
	// Issue kay aata hai if koi file kafi badi rahi to memory kafi jati hai
	//or memory ka issue aasak ta hai

	content, err := ioutil.ReadFile("Example.txt")
	if err != nil {
		fmt.Println("Error while reading Data of the file", err)
		return
	}
	fmt.Println(string(content))

}
