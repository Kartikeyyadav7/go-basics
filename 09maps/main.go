package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["TS"] = "Typescript"
	languages["PY"] = "Python"

	fmt.Println("The map of languages are ", languages)
	fmt.Println("The long for JS is ", languages["JS"])

	delete(languages, "PY") // deletes the PY  in languages
	fmt.Println("The map of languages are ", languages)

	// For loops in GO

	for key, value := range languages {
		fmt.Printf("For key %v, the value is %v\n", key, value)
	}

	// IN the above syntax we can also use the comma, ok syntax to avoid any value like the key we can put a underscore instead of key and it will be ignored

}
