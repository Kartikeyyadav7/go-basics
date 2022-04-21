package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for course - file

type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"price"`
	Author      *Author
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db
var courses []Course

// middleware | helpers - file

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Hello from API")

	r := mux.NewRouter()

	// Seeding - means we are filling our fake db with fake data

	courses = append(courses, Course{CourseId: "2", CourseName: "Reactjs", CoursePrice: 599, Author: &Author{FullName: "Kartikey Yadav", Website: "kartikeyyadav.vercel.app"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Mern stack", CoursePrice: 1599, Author: &Author{FullName: "Kartikey Yadav", Website: "go.vercel.app"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":5000", r))
}

// controller - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from the most awaited Go API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the courses")
	w.Header().Set("Content-Type", "application/json")
	// This will encode the JSON
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-type", "application/json")

	// grab id from  request
	params := mux.Vars(r)

	//  loop through course , find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if the body is nil
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide some value")
	}

	// what about - {}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate new courseId and convert it to string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// getting the id from the request
	params := mux.Vars(r)

	// loop through the id -> remove ->  add with the same ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // remove the element with that id
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//If id is not found
	json.NewEncoder(w).Encode("No course found with the id provided")
	return

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting all the courses")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted the course")
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the given Id")
	return
}
