package controller3
import(
		"fmt"
		"net/http"
		"httprouter"
	)
func Show(w http.ResponseWriter,r *http.Request,p httprouter.Params){
		fmt.Fprintf(w,"Retrive Response")
}
func Update(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	
	fmt.Fprintf(w,"Update Response")

}