package hangmanweb

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

func Victory(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "victory")
}

func Lose(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "lose")
}

func Input(w http.ResponseWriter, r *http.Request) {
	// Retrieve the guessed letter
	guessedLetter := r.FormValue("LettreARecuperer")
	// if guessedLetter == "" {
	// 	http.Error(w, "No letter provided", http.StatusBadRequest)
	// 	return
	// }

	Data.LettreUsed = append(Data.LettreUsed, guessedLetter)

	// Check if the guessed letter exists in the word
	for i, char := range Data.Word {
		if string(char) == guessedLetter {
			// Reveal the guessed letter in the hidden word
			Data.TabHidden[2*i] = guessedLetter
			nothere = false
		}
	}
	// if strings.Join(Data.TabHidden, "") == Data.Word {
	// 	win = true
	// 	http.Redirect(w, r, "/victory", http.StatusSeeOther)
	// 	return
	// }
	if nothere {
		Data.Try--
	}
	win = true
	for _, i := range Data.TabHidden {
		if i == "_" { // If any element is "_", the game is not won
			win = false
			break
		}
	}
	if win {
		Victory(w, r)
	} else if Data.Try == 0 {
		Lose(w, r)
	} else {
		Home(w, r) // Redirect or render the main view
	}
}
