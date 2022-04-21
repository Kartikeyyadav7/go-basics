package main

import "fmt"

func main() {
	fmt.Println("Structs")
	// NO inheritance, no parent or super concept in go

	kartikey := User{"Kartikey", 18, true, "kartikey@go.dev"}
	fmt.Println(kartikey)

	fmt.Printf("The details of Kartikey %+v\n", kartikey)
	fmt.Printf("The name is %v, and the email is %v", kartikey.Name, kartikey.Email)
}

type User struct {
	Name   string
	Age    int
	status bool
	Email  string
}
