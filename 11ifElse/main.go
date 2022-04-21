package main

import "fmt"

func main() {
	fmt.Println("If Else")

	loginCount := 25

	var result string

	if loginCount > 10 {
		result = "Keeps coming back"
	} else if loginCount < 10 {
		result = "Doesn't come at all"
	} else {
		result = "What is this , why do you come exactly 10 times"
	}

	fmt.Println("The result is ", result)

	// Another syntax

	if num := 5; num > 10 {
		fmt.Println("The printed number is greater than 10")
	} else {
		fmt.Println("The printed number is less than 10")
	}

	// if err != nil {

	// }
}
