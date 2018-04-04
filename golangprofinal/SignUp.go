package main

import (
	
	
	"github.com/golangprofinal/DatabaseDB"
	"log"
	"github.com/golangprofinal/routes"
	_ "github.com/golangprofinal/mysql-master" // MySQL Database driver
)


func main() {

	log.Println("Server started on: http://localhost:9000")
	

	DatabaseDB.InitDB("mysql","root","chirag","UserDB")
	routes.Route()

	
}
