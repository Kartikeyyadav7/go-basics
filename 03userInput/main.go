package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome user"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rating for our pizza")

	// comma, ok
	// I can just store the value of input and if error occurs inside of the input variable and err variable
	// One more thing that we can do is if we don't want to do or you don't care about any thig with the error we can just put _ in place of err
	// SO we can use _ for anything in go if we don't want to do anything with that value
	input, _ := reader.ReadString('\n') // We have \n here to denote that we want to read till user put a new line or a enter
	// input , err := reader.ReadString("\n")
	fmt.Println("Thank you for your rating", input)
	fmt.Printf("Type of the rating is %T ", input)
}
