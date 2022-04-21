package main

import "fmt"

func main() {
	//* when there are multiple defer statements we basically use the last in first out appraoch to see which defer will be printed first
	defer fmt.Println("World") // defer does the same thing as the defer in JS it just takes that line of code and simply places it at the bottom of the function just above the last curly braces
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello") // the first thing that will be printed  is Hello after that LIFO - so two is the last in and the first out then one andthen world
	myDefer()            // this will be executed after hello and then 43210 (because the loop print is defered )  and then two , one and world
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
