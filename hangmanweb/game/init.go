package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

const port = ":8080"

type DataForm struct {
	LettreUsed []string
	Words      []string
	NameFill   string
	Word       string
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
func Init() {
	randomWord()
}
func Web(){
	http.HandleFunc("/", Home)
	http.HandleFunc("/input", Input)
	fmt.Println("(http://localhost:8080) - server started on port", port)
	http.ListenAndServe(port, nil)
	fs := http.FileServer(http.Dir("serv/"))
	http.Handle("serv/", http.StripPrefix("serv/", fs))
}
