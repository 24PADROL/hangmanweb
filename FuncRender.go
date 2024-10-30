package main

import (
	"net/http"
	"text/template"
)
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./serv/" + tmpl + ".tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil) 
}
func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contact")
}
