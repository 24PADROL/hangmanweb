package hangmanweb

import (
	// "fmt"
	"net/http"
	"text/template"
	// "fmt"
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

func Victory(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "victory")
}

func Input(w http.ResponseWriter, r *http.Request) {
	// Retrieve the guessed letter
	guessedLetter := r.FormValue("LettreARecuperer")
	// if guessedLetter == "" {
	// 	http.Error(w, "No letter provided", http.StatusBadRequest)
	// 	return
	// }

	// Append the guessed letter to the list of used letters
	Data.LettreUsed = append(Data.LettreUsed, guessedLetter)

	// Check if the guessed letter exists in the word
	for i, char := range Data.Word {
		if string(char) == guessedLetter {
			// Reveal the guessed letter in the hidden word
			Data.TabHidden[2*i] = guessedLetter
		}
	}
	win := true // Assume win initially
	for _, i := range Data.TabHidden {
		if i == "_" { // If any element is "_", the game is not won
			win = false
			break
		}
	}
	if win {
		// fmt.Println("GG, vous avez gagné !")
		WebVictory()
	}
	Home(w, r) // Redirect or render the main view
}
