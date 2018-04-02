package main

import (
	
	
	"github.com/golangpro/DatabaseDB"
	"log"
	"github.com/golangpro/routes"
	_ "github.com/golangpro/mysql-master" // MySQL Database driver
)


func main() {

	log.Println("Server started on: http://localhost:9000")
	

	DatabaseDB.InitDB("mysql","root","chirag","UserDB")
	routes.Route()

	
}
