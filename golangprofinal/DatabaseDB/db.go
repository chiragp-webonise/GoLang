package DatabaseDB
import (
	"database/sql")

var driver,user,pass,name string
func InitDB(dbDriver string,dbUser string,dbPass string,dbName string){


	driver = dbDriver 
	user = dbUser      
	pass = dbPass  
	name = dbName
	
	
}
func Dbconnect() (db *sql.DB){
	db, err := sql.Open(driver, user+":"+pass+"@/"+name)

	
	if err != nil {
		panic(err.Error())
	}

	
	return db
}