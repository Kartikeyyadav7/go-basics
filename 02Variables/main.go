package main

import "fmt"

//* We can't use the no var style outside of any method hello := "bye" is not allowed

//* Though we can declare variables outside using the var keyword
const LoginKey = "asdlkjafeijaldkjfa" //* This is a public constant variable because we have used captial letter to start the variable name
var jwtToken = "halsdkjfaioeja"

const lootlo = "heyheyhey"

func main() {
	var username string = "Kartikey"
	fmt.Println(username)
	fmt.Printf("The type of the variable is %T \n", username)

	var age int = 20
	fmt.Println(age)
	fmt.Printf("The type of the variable is %T \n", age)

	var noOfUsers uint8 = 255
	fmt.Println(noOfUsers)
	fmt.Printf("The type of the variable is %T \n", noOfUsers)

	var someDecimalNo float32 = 2.5754986451
	fmt.Println(someDecimalNo)
	fmt.Printf("The type of the variable is %T \n", someDecimalNo)

	var isMarried bool = false
	fmt.Println(isMarried)
	fmt.Printf("The type of the variable is %T \n", isMarried)

	//* We can also define variables without explicitly telling the type
	var newUser = "Sanskar"
	fmt.Println(newUser)
	fmt.Printf("The type of the variable is %T \n", newUser)

	//* No var style

	newAge := 33
	fmt.Println(newAge)
	fmt.Printf("The type of the variable is %T \n", newAge)

	fmt.Println(LoginKey)
	fmt.Printf("The type of the variable is %T \n", LoginKey)

	fmt.Println(jwtToken)
	fmt.Println(lootlo)

}
