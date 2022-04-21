package main

import "fmt"

func main() {
	fmt.Println("Structs")
	// NO inheritance, no parent or super concept in go

	kartikey := User{"Kartikey", 18, true, "kartikey@go.dev"}
	fmt.Println(kartikey)

	fmt.Printf("The details of Kartikey %+v\n", kartikey)
	fmt.Printf("The name is %v, and the email is %v\n", kartikey.Name, kartikey.Email)

	kartikey.Greet()
	kartikey.newMail()
	fmt.Printf("The details of Kartikey %+v\n", kartikey)
}

type User struct {
	Name   string
	Age    int
	status bool // if there is a smaller case letter that means we don't want it to be avaible publicly except in methods
	Email  string
}

// This is a method and if the first Letter is capital than that means it is public
func (u User) Greet() {
	fmt.Println("Is the user active", u.status)
}

// This function is changing the user email when we are priting it and defining it inside this function but when we are checking the struct value again then it is the same because we are making a copy of that struct we are not passing a reference to that variable
func (u User) newMail() {
	u.Email = "test@go.dev"
	fmt.Println("The new mail is", u.Email)
}
