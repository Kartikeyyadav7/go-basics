package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// const myGetUrl string = "http://localhost:8000"
// const myPostUrl string = "http://localhost:8000/post"
const myPostFormUrl string = "http://localhost:8000/postform"

func main() {
	fmt.Println("Hello from the Web Request")
	// PerformGetRequest(myGetUrl)
	// PerformPostRequest(myPostUrl)
	PerformPostFormRequest(myPostFormUrl)
}

func PerformGetRequest(myUrl string) {
	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("The status code for the response is", response.Status)
	fmt.Println("The content length of the response is", response.ContentLength)

	defer response.Body.Close()

	// We are here using the strings builder method to create a string and then storing the content inside of that responseString
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("The byte count is", byteCount)
	fmt.Println(responseString.String())

	//This is one way to handle the content , another way is above
	// fmt.Println("The content is ", string(content))
}

func PerformPostRequest(myUrl string) {
	requestBody := strings.NewReader(`
		{
			"name" : "Kartikey",
			"age" : 21,
			"gender" : "male"
		}	
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("The response is", response)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("The content is", string(content))
}

func PerformPostFormRequest(myUrl string) {
	data := url.Values{}

	data.Add("firstName", "Kartikey")
	data.Add("lastName", "yadav")
	data.Add("age", "21")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("The content is", string(content))

}
