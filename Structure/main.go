package main

import "fmt"

//Structer store different data type variable in single variable
// int,float,bool.char,person,contact

type Person struct {
	First_Name string
	Last_Name  string
	age        int
}

type contact struct {
	Email string
	phone string
}

type address struct {
	House int
	area  string
	state string
}

type Employee struct {
	Person_info    Person
	Person_contact contact
	Person_address address ``
}

func main() {
	fmt.Println("Learning Structer in GOLANG")
	var Pritam Person

	fmt.Println("Pritam Person:", Pritam)

	Pritam.First_Name = "Pritam"
	Pritam.Last_Name = "Bhure"
	Pritam.age = 24

	fmt.Println("Pritam Info:", Pritam)
	fmt.Println("")

	// Method 2
	Britnu := Person{
		First_Name: "Britnu",
		Last_Name:  "Bhure",
		age:        3,
	}
	fmt.Println("Britnu Info:", Britnu)
	fmt.Println("")

	// Method 3
	var arzoo = new(Person)
	arzoo.First_Name = "arzoo"
	arzoo.Last_Name = "Bhure"
	arzoo.age = 20
	fmt.Println("Arzoo Info:", arzoo)

	var employee1 Employee

	employee1.Person_info = Person{
		First_Name: "Pritam",
		Last_Name:  "Bhure",
		age:        24,
	}

	employee1.Person_contact = contact{
		Email: "Pritambhure558@gmail.com",
	}
	// also we can write this
	employee1.Person_contact.phone = "8080766512"

	employee1.Person_address = address{
		House: 401,
		state: "Maharastra",
		area:  "Kazi nagar bhandara",
	}

	fmt.Println("Info of the employee", employee1)
}
