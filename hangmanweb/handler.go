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

func Menu(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "menu")
}

func letterAlreadyGuessed(s string) bool {
	for _, i := range Data.LettreUsed {
		if s == i {
			return true
		}
	}
	return false
}

func formateLetter(guessedLetter string) string {
	if guessedLetter == "é" || guessedLetter == "è" || guessedLetter == "ë" || guessedLetter == "ê" {
		guessedLetter = "e"
	}
	if guessedLetter == "à" || guessedLetter == "â" {
		guessedLetter = "a"
	}
	if guessedLetter == "ù" || guessedLetter == "û" {
		guessedLetter = "u"
	}
	if guessedLetter == "î" || guessedLetter == "ï" {
		guessedLetter = "i"
	}
	if guessedLetter == "ô" || guessedLetter == "ö" {
		guessedLetter = "o"
	}
	if guessedLetter == "ç" {
		guessedLetter = "c"
	}
	return guessedLetter
}

func Input(w http.ResponseWriter, r *http.Request) {
	// Retrieve the guessed letter
	guessedLetter := r.FormValue("LettreARecuperer")
	guessedLetter = formateLetter(guessedLetter)
	if guessedLetter == Data.Word {
		Victory(w, r)
	} else {
		if !letterAlreadyGuessed(guessedLetter) {
			Data.LettreUsed = append(Data.LettreUsed, guessedLetter)
			for i, char := range Data.Word {
				if string(char) == guessedLetter {
					// Reveal the guessed letter in the hidden word
					Data.TabHidden[2*i] = guessedLetter
					nothere = false
				}
			}
			if len(guessedLetter) > 1 {
				Data.Try--
			}
			if nothere {

				Data.Try--
			}
			nothere = true
			win = true
			for _, i := range Data.TabHidden {
				if i == "_" { // If any element is "_", the game is not won
					win = false
					break
				}
			}
		}
		if win {
			Victory(w, r)
		} else if Data.Try <= 0 {
			Lose(w, r)
		} else {
			Home(w, r) // Redirect or render the main view
		}
	}
}
