package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var Url string = "https://www.aviatecoders.com"

func main() {
	fmt.Println("Hello from web requests")
	response, err := http.Get(Url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The response type is %T\n ", response)

	defer response.Body.Close() //* We have to manually close the connection

	dataByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(nil)
	}

	content := string(dataByte)

	fmt.Println("Content is", content)

}
