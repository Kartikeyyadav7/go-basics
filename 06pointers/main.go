package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	var ptr *int
	fmt.Println("The value of just a pointer is ", ptr)

	myNumber := 55
	var numPtr *int = &myNumber

	//as the type is infered automatically we can just write like var numPtr = &myNumber

	fmt.Println("The pointer to my Numbers is ", numPtr)
	fmt.Println("The pointer to my Numbers is ", *numPtr)

	*numPtr = *numPtr + 2
	fmt.Println("The numbers pointer is increased by", myNumber)

}
