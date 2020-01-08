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

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/fav", fav)
	http.HandleFunc("/about", about)
   	http.ListenAndServe(":8088", nil)
}
