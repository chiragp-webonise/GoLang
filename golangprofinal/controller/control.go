package controller

import (
	"database/sql"
	"net/http"
	"html/template" 
	"log"
	"regexp"
	_"github.com/golangpro/mysql-master"
	"github.com/golangpro/DatabaseDB"
	"github.com/golangpro/model"
	)

var tmpl = template.Must(template.ParseGlob("view/*.html"))
type errorS struct{
	Errormsg1 string
	Errormsg2 string
    Errormsg3 string
    Errormsg4 string
 }
type Users struct{
	FirstName string
	LastName string
	Email string
} 
var db *sql.DB

func Insert(w http.ResponseWriter, r *http.Request) {

	db=DatabaseDB.Dbconnect()
	e:=errorS{}
	flag:=1
	
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	passRegexp := regexp.MustCompile("^[A-Za-z]\\w{6,}[A-Za-z]$")
	
	if r.Method == "POST" {

	
		fname:= r.FormValue("fname")
		lname:= r.FormValue("lname")
		email:= r.FormValue("email")
		psw := r.FormValue("psw")
		pswRepeat := r.FormValue("pswRepeat")
		
		if !emailRegexp.MatchString(email) {

			flag=0
			e.Errormsg2="Invalid email"
			log.Println(e.Errormsg2)
			
		}else{
			e.Errormsg2=""
		}
		if !passRegexp.MatchString(psw) {

			flag=0
			e.Errormsg3="Invalid password, the password must be at least 8 characters long and start and end with a letter."
			
			log.Println(e.Errormsg3)
		}else{
			e.Errormsg3=""
		}		
		 if psw != pswRepeat {

			flag=0
			e.Errormsg1="Repeat password did not match"
			log.Println(e.Errormsg1)
			
		}else{
			e.Errormsg1=""
		}
		if flag==1{
			
			var email,psw, eid string
			if r.Method == "POST" {

				
				email = r.FormValue("email")
			}
			
			selDB, err := db.Query("SELECT email FROM Users WHERE email=?",email)
			if err != nil {
				panic(err.Error())
			}

			
			for selDB.Next() {
				
				

				
				err = selDB.Scan(&eid)
				if err != nil {
					panic(err.Error())
				}	
			}
			if eid==""{
						model.Create(db)
						psw = r.FormValue("psw")
						model.Insert(db,fname,lname,email,psw)
						
						tmpl.ExecuteTemplate(w, "Login.html",nil)
						log.Println("INSERT: Name: " + fname + " | E-mail: " + email)

				}else{
					e.Errormsg4="Email already exist"
					tmpl.ExecuteTemplate(w, "Registration.html",e)

					}
		}else{
			tmpl.ExecuteTemplate(w, "Registration.html",e)
	}
	
	defer db.Close()
	
	
}
}
func Registration(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Registration.html",nil)
}
func Login(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Login.html",nil)
}
func Check(w http.ResponseWriter, r *http.Request) {
	
	db=DatabaseDB.Dbconnect()
	e:= errorS{}
	
	var email,psw,pass, eid string
	if r.Method == "GET" {

		
		email = r.FormValue("email")
		psw   = r.FormValue("psw")
	}
	
	eid,pass=model.LoginCheck(db,email)

		if (email==eid && psw==pass){
			var fname, lname, eid string
			
				u := Users{}

					fname,lname,eid=model.DashboardData(db,email)
					u.FirstName = fname
					u.LastName = lname
					u.Email = eid
				
			tmpl.ExecuteTemplate(w, "DashBoard.html",u)			
		}else{		
			e.Errormsg4="email or password invalid"
			tmpl.ExecuteTemplate(w, "Login.html",e)	
		}

	defer db.Close()

}