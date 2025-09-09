package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	rai := User{"Rai", "raisumant@go.dev", true, 16}
	fmt.Println(rai)
	fmt.Printf("rai details are: %+v\n", rai)
	fmt.Printf("Name is %v and email is %v.\n", rai.Name, rai.Email)
	rai.GetStatus()
	rai.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", rai.Name, rai.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}