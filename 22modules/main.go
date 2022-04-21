package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Modules in GO")

	//* The use of go mod init <your github username> command is that it helps some other third party to know what go version I am using

	/* Now I have downloaded a package from the internet and it will show on the go.mod file along with the version and a comment saying that
	it is indirect because we are not using that package,but as soon as we start using it and run the go mod tidy command it will remove that
	comment. As I downloaded the package a go.sum file is created which actually contains the hash of the downlaoded package

	One more thing whenever we download a package it is not stored inside of the current working directory , it gets stored where the go file is
	downloaded , and whenever we want to request that package again it goes and checks the available cache and if it is present it fetches from
	there

	Go mod tidy command is an expensive command but it will tidy up all the packages that you are using , that means it will update their statuses, remove
	those packages that you are not using and does all the good things

	Go mod verify is a expensive command that will just change verify the go sum file with the things on the internet so that hash is proper

	Go list will give me the list of the packages my application is dependent on - you can use some flags like go list -m all which will list all the proper dependencies

	Go mod why github.com/gorilla/mux is the command that will tell us why we are using that library and which files are dependent on that

	Go mod graph is the command that will tell us the files that are dependent on each other

	Now when ever you download some packages from the internet it goes to the cache file so what if you want that package in your file where you are
	working , for that reason you can include it using go mod vendor - which will bring up the vendor folder that will contain all the dependencies
	but you will need to use the go run -mod=vendor main.go that will run the command from the vendor file or normally if you use go run main.go
	it will go the cache and run the file

	*/
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hey there mod user")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Go</h1>"))
}
