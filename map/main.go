package main

import "fmt"

//map is similler to hash maping (Key value pair) store in unorder way

func main() {

	// declaration of map

	studentsGrades := make(map[string]int)

	studentsGrades["Pritam"] = 80
	studentsGrades["Bhavesh"] = 99
	studentsGrades["arzoo"] = 35
	studentsGrades["jivika"] = 90

	fmt.Println("Marks of students", studentsGrades["Bhavesh"])

	fmt.Println("Marks of students pritam is", studentsGrades["Pritam"])
	//studentsGrades["Pritam"] = 100
	//fmt.Println("The new Marks of students pritam is", studentsGrades["Pritam"])

	//delete data take to values(map and key)
	delete(studentsGrades, "Pritam")

	// checking key is present or not in the map
	grades, Exists := studentsGrades["Britnu"] //returns two values (values & boolen value (present or not))
	fmt.Println("The grades od the Britnu is", grades)
	fmt.Println("Britnu is present or not in the map", Exists)

	info := make(map[string]string)
	var key, value string

	counter := 0
	for {
		fmt.Println("Enter key ")
		fmt.Scan(&key)
		counter++
		if key == "done" {

			break
		}

		fmt.Println("Enter Value ")
		fmt.Scan(&value)

		info[key] = value

	}
	for key, values := range info {
		fmt.Printf("The key is %s and the values is %s\n", key, values)
	}

}
