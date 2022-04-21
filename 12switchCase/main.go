package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello from switch case")

	rand.Seed(time.Now().Unix())
	randomNum := rand.Intn(6) + 1

	fmt.Println("The random no. is", randomNum)

	switch randomNum {
	case 1:
		fmt.Println("You can open and move in 1 ")
	case 2:
		fmt.Println(" move to 2 places ")
	case 3:
		fmt.Println("move to 3 places")
	case 4:
		fmt.Println("move to 4 places ")
		fallthrough // when we say fallthrough it will print 4 as well as 5
	case 5:
		fmt.Println("move to 5 places ")
	case 6:
		fmt.Println("move to 6 places and roll the dice again")
	default:
		fmt.Println("What the hell was that !!!")
	}
}
