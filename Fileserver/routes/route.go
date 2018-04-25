package routes

import (
"net/http"
"github.com/Fileserver/controller"
)


func Route(){


	// URL management
	// Manage templates
	http.HandleFunc("/", controller.Downloadfile)    // INDEX :: Show all registers
	//http.HandleFunc("/show", Show) // SHOW  :: Show only one register
	//http.HandleFunc("/new", New)   // NEW   :: Form to create new register
	//http.HandleFunc("/edit", Edit) // EDIT  :: Form to edit register
	http.HandleFunc("/download", controller.Download)
	// Manage actions
	
	//http.HandleFunc("/update", Update) // UPDATE :: Update register
	//http.HandleFunc("/delete", Delete) // DELETE :: Destroy register

	// Start the server on port 9000
	//panic(http.ListenAndServe(":9000", http.FileServer(http.Dir("/github.com/golangpro/Static"))))
	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	http.ListenAndServe(":9000",nil)

}