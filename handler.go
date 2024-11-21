package main

import (
	// "fmt"
	"net/http"
	"text/template"
)

type DataForm struct { 
	LettreUsed []string
}

var Data DataForm

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./serv/" + tmpl + ".tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, Data)
}

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home")
}

func Input(w http.ResponseWriter, r *http.Request) {
	Data.LettreUsed = append(Data.LettreUsed, r.FormValue("LettreARecuperer"))
	Home(w, r)
}
