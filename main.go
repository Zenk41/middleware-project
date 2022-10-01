package main

import (
	"middleware-project/database"
	"middleware-project/routes"

)


func main() {
	// initialize database
	database.InitDB()
	// auto migrate table and create table if not exist 
	database.InitMigrate()
	// initialize routes
	e := routes.New()
	// port
	e.Logger.Fatal(e.Start(":8000"))
}
