package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	
)

//structure to get and insert the data in database
type Movie struct{
	MovieID string `json:"movieid"`
	MovieName string `json:"moivename"`
}
//structure to get the json responce 
type JsonResponse struct{
	Type 	string	`json:"type"`
	Data	[]Movie `json:"data`
	Message string `json:"message"`
}



//constants for database services
const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "abc"
	dbname = "userinfo"
)



// Function for handling errors
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}



func setupDB() *sql.DB{
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",user,password,dbname)
	db,err := sql.Open("postgres", dbinfo)
	if err != nil{
		panic(err)
	}
	return db
}

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/info",getalldata).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000",r))
}


//controllers 
func getalldata(w http.ResponseWriter, r *http.Request)  {
	db := setupDB()

	rows , err := db.Query("select * from moives")
	checkErr(err)

	// var response []JsonResponse
    var movies []Movie

    // Foreach movie
    for rows.Next() {
        var id int
        var movieID string
        var movieName string

        err = rows.Scan(&id, &movieID, &movieName)

        // check errors
        checkErr(err)

        movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
    }

    var response = JsonResponse{Type: "success", Data: movies}

    json.NewEncoder(w).Encode(response)
}

