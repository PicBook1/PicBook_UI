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
func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
   	http.ListenAndServe(":8088", nil)
}
