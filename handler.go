package main

import (
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, html string) {
	t, err := template.ParseFiles("./serv/" + html + ".html")
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
