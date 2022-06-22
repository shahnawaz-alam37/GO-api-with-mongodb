package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shahnawaz-alam37/newrepo/router"
)

func main()  {
	fmt.Println("mongoDB API")
	r := router.Router()
	fmt.Println("live")
	log.Fatal(http.ListenAndServe(":5000",r))
	fmt.Println("server at local host port 5000")
}