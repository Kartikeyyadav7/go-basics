package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello from files")
	content := "Welcome to Aviate Coders, where you can learn about a lot things regarding ... you know what "

	file, err := os.Create("./aviate.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("The length of the file is ", length)
	defer file.Close()
	readFile("E:\\godev\\17files\\aviate.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("The data in the file is ", dataByte)         // this is the array buffer
	fmt.Println("The data in the file is ", string(dataByte)) // this is the proper data

}

// As the if statement for checking error is being used too many times people generally make a function and call that in place of iferr statement
// func checkNilErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
