package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)
type Course struct{
	CourseId string `json:"course Id"`
	CourseName string `json:"course name"`
	CoursePrice int `json:"price"`
	Author 		*Author `json:"author"`
}
type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}
//fake database 
var courses []Course

//middleware or helpre - file
func (c *Course) Isempty() bool{
	return c.CourseId == "" && c.CourseName == ""
}


func main(){
	fmt.Println("simple api")
}

//controllers
//serve home route
func serverHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>welcome to my first api</h1>"))
}

func getallcourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOnecourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("get one course")
	w.Header().Set("Content-Type","application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses,find the matching id and return the responce
	for _,course := range courses{
		if course.CourseId==params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}