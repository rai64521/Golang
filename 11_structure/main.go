package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	sumant := User{"sumant", "sumant@go.dev", true, 20}
	fmt.Println(sumant)
	fmt.Printf("sumant rai details are: %+v\n", sumant)
	fmt.Printf("Name is %v and email is %v.", sumant.Name, sumant.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}