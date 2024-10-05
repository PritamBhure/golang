package main

import (
	"fmt"
	"time"
)

func main() {

	current_date := time.Now()
	fmt.Println("Current Time", current_date)
	// 2006(yyyy)-01(mm)-02(dd) Time =  15(24 HH)/03(12 HH) :01(MM):04(SS) PM monday(Current Day)
	formatted_time := current_date.Format("01-02-2006 03:04:03 PM Monday")

	fmt.Println("Current Date", formatted_time)

	time_Layout := "01-02-2006"
	string1 := "07-23-2050"

	formatted_time2, _ := time.Parse(time_Layout, string1)
	fmt.Println("Convert string into date", formatted_time2)

	//manupalation with day

	new_date := current_date.Add(24 * time.Hour)
	fmt.Println("New Date", new_date)

	new_date5 := new_date.Format("01-02-2006 03:04:03 PM Monday")

	println("new add date")
	println("The Event date", new_date5)
}
