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
func fav(w http.ResponseWriter, r *http.Request){
  templ.ExecuteTemplate(w,"fav.layout",nil)
 }
func about(w http.ResponseWriter, r *http.Request){
  templ.ExecuteTemplate(w,"about.layout",nil)
 }

func login(w http.ResponseWriter, r *http.Request){
	templ.ExecuteTemplate(w,"login.layout",nil)
 }
