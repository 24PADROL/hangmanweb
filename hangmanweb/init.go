package hangmanweb

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

type DataForm struct {
	LettreUsed []string
	Words      []string
	NameFill   string
	Word       string
	TabHidden  []string
	Letter     string
	Try        int
}

var Data DataForm

var win bool = false

var nothere bool = true

var nameFill string = "motsimple.txt"

const port = ":8080"

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
	Data.Word = strings.ToLower(Data.Words[rand.Intn(200)])
}

func printHidden() {
	for i := 0; i < len([]rune(Data.Word)); i++ {
		Data.TabHidden = append(Data.TabHidden, "_")
		Data.TabHidden = append(Data.TabHidden, " ")
	}
}

func Init() {
	Data = DataForm{}
	Data.Try = 10
	randomWord()
	printHidden()
}

func Reset(w http.ResponseWriter, r *http.Request) {
	Init()                                            // Réinitialise les données de jeu
	http.Redirect(w, r, "/home", http.StatusSeeOther) // Redirige vers la page d'accueil
}

func Web() {
	http.HandleFunc("/", Menu)     // Menu page
	http.HandleFunc("/home", Home) // Game page
	http.HandleFunc("/victory", Victory)
	http.HandleFunc("/lose", Lose)
	http.HandleFunc("/input", Input)
	http.HandleFunc("/reset", Reset)
	http.HandleFunc("/thankyou", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./serv/thankyou.html")  })

	fs := http.FileServer(http.Dir("serv/"))
	http.Handle("/serv/", http.StripPrefix("/serv/", fs))

	fmt.Println("(http://localhost:8080) - server started on port", port)
	http.ListenAndServe(port, nil)

}
