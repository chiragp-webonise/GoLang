package controller

import (
	"net/http"
	"html/template" 
	"os"
	"fmt"
	"io/ioutil"
    "bytes"
	)
var tmpl = template.Must(template.ParseGlob("view/*.html"))
func Downloadfile(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Download.html",nil)
}


func Download(w http.ResponseWriter, r *http.Request) {

          Url:= r.FormValue("url")
          streamPDFbytes, err := ioutil.ReadFile(Url)

          if err != nil {
                  fmt.Println(err)
                  os.Exit(1)
          }

          b := bytes.NewBuffer(streamPDFbytes)


          
          w.Header().Set("Content-Disposition","attachment;filename="+Url)
          
          if _, err := b.WriteTo(w); err != nil { 
                  fmt.Fprintf(w, "%s", err)
          }

          w.Write([]byte("Download complete"))
  }