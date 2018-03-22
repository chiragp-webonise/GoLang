package main

import(
		"flag"
		//"fmt"
		"log"
		"net/http"
		"httprouter"
		"github.com/restapi/controller3"
	)
var(
	addr=flag.String("addr",":8081","http service address")
)
func main() {
	flag.Parse()
	r:=httprouter.New()//router instance
	r.GET("/retrive",controller3.Show)
	r.PUT("/update",controller3.Update)
	err:= http.ListenAndServe(*addr,r)
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}
