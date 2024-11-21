package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"text/template"
	"math/rand"
)

type DataForm struct {
	LettreUsed []string
	Words []string
	NameFill string
	Word 	string
}

var Data DataForm

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
		Data.Words = append(Data.Words, scanner.Text())
	}
	if len(Data.Words) == 0 {
		fmt.Println("le fichier ne contient rien")
		return
	}
	Data.Word = Data.Words[rand.Intn(200)]
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
