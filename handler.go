package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type DataForm struct {
	LettreUsed []string
}

var Data DataForm
var words []string
var nameFill string = "motsimple.txt"

func randomWord() {
	fichier, err := os.Open(nameFill)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	defer fichier.Close()
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if len(words) == 0 {
		fmt.Println("le fichier ne contient rien")
		return
	}
}

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
