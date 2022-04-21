package main

import "fmt"

func main() {
	fmt.Println("Arrays in go")

	// This is how we define arrays in GO
	var fruitlist [4]string

	fruitlist[0] = "mango"
	fruitlist[1] = "banana"
	fruitlist[3] = "oranges"

	fmt.Println("The fruitList is ", fruitlist)               // If you see the output clearly you can see that there is space between banana and oranges for a element that we have missed
	fmt.Println("The length of fruitlist is", len(fruitlist)) //The len method returns the number 4 even if there are only 3 elements present in the array

	var vegList = [3]string{"ladyfinger", "beans", "potato"}

	fmt.Println("The list of vegetables are ", vegList)

}
