package main

import(
		"flag"
		//"fmt"
		"log"
		"net/http"
		"httprouter"
		"github.com/restapi/controller2"
	)
var(
	addr=flag.String("addr",":8081","http service address")
	Data map[string]string
)
func main() {
	flag.Parse()
	Data=map[string]string{}
	r:=httprouter.New()//router instance
	r.GET("/entry/:key",controller3.Show)
	r.GET("/list",controller.Show)
	r.PUT("/entry/:key/:value",controller2.Update)
	err:= http.ListenAndServe(*addr,r)
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}
/*var data map[string]string
func Show(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	k:=p.ByName("key")
	if k == ""{
		fmt.Fprintf(w,"Read list:%v",data)
		return
	}
	fmt.Fprintf(w,"Read entry: data[%s]=%s",k,data[k])
}
func Update(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	k:= p.ByName("key")
	v:= p.ByName("value")

	data[k]= v
	fmt.Fprintf(w,"Updated: data[%s]=%s",k,data[k])

}*/