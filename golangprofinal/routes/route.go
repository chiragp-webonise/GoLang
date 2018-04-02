package routes

import (
"net/http"
"github.com/golangpro/controller"
)


func Route(){


	
	http.HandleFunc("/", controller.Login)    
	
	http.HandleFunc("/register", controller.Registration)
	
	http.HandleFunc("/insert", controller.Insert) 
	http.HandleFunc("/check", controller.Check)
	
	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	http.ListenAndServe(":9000",nil)

}