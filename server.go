package main

import (
	"html/template"
	 "net/http"	 
 )

 var templ *template.Template
 func init(){
 	templ = template.Must(template.ParseGlob("assets/templates/*"))
}
func index(w http.ResponseWriter, r *http.Request){
	templ.ExecuteTemplate(w,"index.layout",nil)
}
