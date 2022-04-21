package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://aviatecoders.com:5000/javascript?blogName=closure&loggedIn=true"

func main() {
	fmt.Println("Hello from the URLS")
	fmt.Println(myUrl)
	// Parsing the url
	result, _ := url.Parse(myUrl)

	// fmt.Println("Resulted Url after parsing", result)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The type of query params is %T\n", qparams) // this qparams is a map and I can access the values of this map
	fmt.Println("Query params", qparams["blogName"])

	for _, val := range qparams {
		fmt.Println("Params is", val)
	}

	// Making of a URL

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Path:     "/code",
		Host:     "aviatecoders.com",
		RawQuery: "user=Kartikey",
	}

	anotherUrl := partsOfUrl.String() // here we have used .String we could also would have done like string(partsOfUrl) both are valid
	fmt.Println("The url is", anotherUrl)
}
