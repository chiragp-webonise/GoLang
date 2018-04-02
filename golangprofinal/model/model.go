package model
import (
	"database/sql"
	)

func Insert(db *sql.DB,fname string,lname string,email string,psw string) {

	insForm, err := db.Prepare("INSERT INTO Users(first_name, last_name, email, password) VALUES(?,?,?,?)")
			if err != nil {
				panic(err.Error())
			}

			
			insForm.Exec(fname, lname, email, psw)


}
func Create(db *sql.DB){
	_,err := db.Exec("CREATE TABLE IF NOT EXISTS Users(id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,first_name varchar(30),last_name varchar(30),email varchar(30),password varchar(20) )" )
   if err != nil {
       panic(err)
   }
}
func LoginCheck(db *sql.DB,email string)(string,string){
	var eid,pass string
	selDB, err := db.Query("SELECT email,password FROM Users WHERE email=?",email)
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		
		err = selDB.Scan(&eid,&pass)
		if err != nil {
			panic(err.Error())
		}
	}
		return eid,pass
}
func DashboardData(db *sql.DB,email string)(string,string,string){
	var fname, lname, eid string
	selDB, err := db.Query("SELECT first_name,last_name,email FROM Users WHERE email=?",email)
				if err != nil {
					panic(err.Error())
				}
	for selDB.Next() {
					
					err = selDB.Scan(&fname, &lname, &eid)
					if err != nil {
						panic(err.Error())
					}
				}
				return fname,lname,email			
}