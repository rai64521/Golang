package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://api.github.com/users/rai64521"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "api.github",
		Path:    "/users",
		RawPath: "user=rai6451",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}