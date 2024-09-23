package main

import (
	"chatApp/internal/db"

	_ "github.com/lib/pq"
)

func main() {
	db.ConnectToDb()

	// fmt.println(db.migrator().hastable(&user{}))
	// fmt.println(db.migrator().hascolumn(&user{}, "email"))
	// fmt.println(db.migrator().hascolumn(&user{}, "name"))
	// fmt.println(db.migrator().hascolumn(&user{}, "password"))
	// //  claim := JWT.Claims{
	//    Role: "Boss",
	//  }
	//  JWT.GetToken()
	//  fmt.Println(claim)
}
