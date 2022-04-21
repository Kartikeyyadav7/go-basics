package main

import "fmt"

func main() {
	fmt.Println("Functions")
	greeter()

	result := adder(2, 5)
	fmt.Println("The result is", result)

	secondResult, message := adder2(5, 10)
	fmt.Println("The secondResult is", secondResult, "my message is", message)

	proResult := proAdder(5, 6, 8, 22)
	fmt.Println("The pro result is", proResult)
}

// If a function returns something then we just have to use the int to denote what is the return type of teh function
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// I can also return two values and wrap them in () brackets
func adder2(valOne int, valTwo int) (int, string) {
	return valOne + valTwo, "Ok added"
}

// This function is going to take unknown no. of arguments and add all of them together, the ...int means that all the values will be integer
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func greeter() {
	fmt.Println("Greeting from Greeter")
}

// There exists function without name and IIFE as well
