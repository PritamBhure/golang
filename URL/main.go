package main

import (
	"fmt"
	"net/url"
)

func main() {
	myURL := "https://www.youtube.com/watch?v=NTpbbQUBbuo"
	//url.Parse() is use convert string into URL
	parseurl, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error while converting string into URL", err)
		return
	}

	fmt.Printf("Type of URL: %T\n", parseurl)
	fmt.Println("Schem of URl", parseurl.Scheme)
	fmt.Println("Host of URL:", parseurl.Host)
	fmt.Println("Path of URL: ", parseurl.Path)
	fmt.Println("RawQuery of URL: ", parseurl.RawQuery)

}
