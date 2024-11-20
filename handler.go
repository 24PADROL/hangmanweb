package main

import (
	// "fmt"
	"net/http"
	"text/template"
)

type DataForm struct { 
	Lettre string
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
	// switch r.Method {
	// case "GET":
	// 	http.ServeFile(w, r, "home.tmpl")
	// case "POST":
	// 	if err := r.ParseForm(); err != nil {
	// 		fmt.Fprintf(w, "ParseForm() err : %v", err)
	// 		return
	// 	}
	// 	name := r.FormValue("name")

	// 	fmt.Fprintf(w, "lettre = %s\n", name)
	// }
}

func Input(w http.ResponseWriter, r *http.Request) {
	Data.Lettre = r.FormValue("LettreARecuperer")
	
	Home(w, r)
}