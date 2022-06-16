package main

import (
	"database/sql"
	"fmt"
)

const (
	user = "postgres"
	password = "abc"
	dbname = "movies"
)

func main()  {
	connecionstring := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",user,password,dbname)
	db,err := sql.Open("postgres",connecionstring)

	if err != nil {
		panic(err)
	}

	quri,err := db.Query("select *from movies")

	if err != nil {
		panic(err)
	}
	fmt.Println(quri)
	db.Close()
}