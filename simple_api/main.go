package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author,omitempty"`
}

type Author struct {
	FullName string `json:"fullName"`
	WebSite  string `json:"website"`
}

//fake Db
var courses []Course

var portNumber = ":8000"

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("GOLANG API...")
	r := mux.NewRouter()

	//create a fake courses data
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "2021 Road to Django Developer",
		CoursePrice: 299,
		Author: &Author{
			FullName: "Keshav R Karanth",
			WebSite:  "keshavrkaranth.me",
		},
	})
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "2021 Road to Go lang Developer",
		CoursePrice: 1399,
		Author: &Author{
			FullName: "Darshan N Shetty",
			WebSite:  "Iamdarshan.com",
		},
	})
	// routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getSingleCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(portNumber, r))

}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home page from GO-API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Course by ID")
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No Course Found with that ID")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a single course")
	w.Header().Set("content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Body should be passed")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Course Name should be passed")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a single course")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with that ID")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a single course")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted CourseFound with that ID")
			break
		}
	}
	json.NewEncoder(w).Encode("No Course Found with that ID")
	return
}
