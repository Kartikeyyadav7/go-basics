package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices ")

	// IN slices we don't have to define the number of elements it is gonna take , we can just use the append function to add new elements
	var fruitList = []string{"Apple", "Banana", "Peach"}

	fruitList = append(fruitList, "Mango", "Oranges")

	fruitList = append(fruitList[1:]) // What happens is when we put 1: we are telling Go to just print the values from index 1 to the end

	// fruitList = append(fruitList[1:3]) // here it is going to print from 1 to 2 because we don't include 3

	fmt.Println("The list of fruits are ", fruitList)

	highScores := make([]int, 4) //We are using the make syntax here to allocate memory with 4 memory spaces for a slice

	highScores[0] = 555
	highScores[1] = 452
	highScores[2] = 324
	highScores[3] = 850
	// highScores[4] = 342 // This line will give an error because first we defined the slice with only 4 memory spaces and then we are now adding new values, so GO is going to complain

	highScores = append(highScores, 757) // this is line is not giving error because when we use append it will reallocate the memory and all of the elements will be reallocated

	fmt.Println(highScores)

	sort.Ints(highScores) // the sort package is only avaiable for slices

	// fmt.Println("The sorted array is", highScores)

	//* How to remove a value from slices based on index

	var courses = []string{"javascript", "go", "swift", "react"}

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...) // this syntax basically does is that first it starts from all the courses and then it ends at index, again it starts with the index + 1 and then ends at the last element and we have use the ... because we going to take the rest of the parameters

	fmt.Println("The new array with the removed index is", courses)

}
