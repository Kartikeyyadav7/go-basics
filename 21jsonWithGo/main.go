package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go")
	// EncodeJson()
	DecodeJson()

}

type course struct {
	Name     string `json:"courseName"` // if we don't define the json here it will called as Name in json format and we don't want that
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // if we use - then it means that , that particular value will be not outputted when the json is called
	Tags     []string `json:"tags,omitempty"` // here we have changed the Tags name to tags and used omitempty so that if we have nil passed then that field will not be shown, also keep in mind that there should not be any space between tags,omitemtpy or else you will get error
}

func EncodeJson() {
	aviateCourses := []course{
		{"Reactjs", 1288, "aviatecoders.com", "abe345", []string{"web-dev", "js"}},
		{"Angularjs", 128, "youtube.com", "e345", nil},
		{"MERN", 14288, "aviatecoders.com", "asldkjfeabe345", []string{"full-stack", "js"}},
	}

	//* package this data as JSON data

	finalJson, err := json.MarshalIndent(aviateCourses, "", "\t") // Here we are using the JSON package , to encode json and in that we are using the Marshal indent to indent out code, and we are using \t to indent all the json data

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"courseName": "Reactjs",
		"Price": 1288,
		"website": "aviatecoders.com",
		"tags": ["web-dev","js"]
	}	
	`)

	var aviateCourses course

	checkValidJson := json.Valid(jsonDataFromWeb)

	if checkValidJson {
		fmt.Println("The json is valid")
		json.Unmarshal(jsonDataFromWeb, &aviateCourses)
		fmt.Printf("%#v\n", aviateCourses)
	} else {
		fmt.Println("THE JSON WAS NOT VALID")
	}

	//* In some cases you want to just add data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("The key is %v and the value is %v and the type of data is %T\n", k, v, v)
	}
}
