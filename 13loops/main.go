package main

import "fmt"

func main() {
	fmt.Println("Loops ")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("The index is %v, and then day is %v\n", index, day)
	}

	rougeValue := 1
	// THis type of for loop is similar to while loop like in other languages
	for rougeValue < 10 {
		//* we can use break in go
		// if rougeValue == 5 {
		// 	break
		// }

		//* we are using goto to go to the label
		if rougeValue == 3 {
			goto aviate
		}

		//* Continue in loops
		// if rougeValue == 5 {
		// 	rougeValue++
		// 	continue
		// }

		fmt.Println("The rouge value is", rougeValue)
		rougeValue++
	}
	// this is label
aviate:
	fmt.Println("Jumping at aviate coders")
}
