package main

import (
	"fmt"
	"net/http"
)
//fake db
type Info struct{
	Name string  `json:"name"`
	Pin string   `json:"pin"`
}
var obj []Info


func main()  {
	fmt.Println("new api")
}

//controllers
func getallinfo(w http.ResponseWriter, r *http.Request)  {
	fmt.println()
}