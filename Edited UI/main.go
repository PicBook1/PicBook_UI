package main

import (
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("UI/templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.layout", nil)
}
func fav(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "favorite.layout", nil)
}
func login(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "login.layout", nil)
}
func about(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.layout", nil)
}

func main() {
	//mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("UI/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/fav", fav)
    http.HandleFunc("/about", about)
    http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}