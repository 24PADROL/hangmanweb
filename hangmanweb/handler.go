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
func image(n int){
	switch n {
	case 1:
		Data.ImagePath = "serv/image/Hangman_0_vies-removebg-preview.png"
	case 2:
		Data.ImagePath = "serv/image/Hangman_1_vies-removebg-preview.png"
	case 3:
		Data.ImagePath = "serv/image/Hangman_2_vies-removebg-preview.png"
	case 4:
		Data.ImagePath = "serv/image/Hangman_3_vies-removebg-preview.png"
	case 5:
		Data.ImagePath = "serv/image/Hangman_4_vies-removebg-preview.png"
	case 6:
		Data.ImagePath = "serv/image/Hangman_5_vies-removebg-preview.png"
	case 7:
		Data.ImagePath = "serv/image/Hangman_6_vies-removebg-preview.png"
	case 8:
		Data.ImagePath = "serv/image/Hangman_7_vies-removebg-preview.png"
	case 9:
		Data.ImagePath = "serv/image/Hangman_8_vies-removebg-preview.png"
	case 10:
		Data.ImagePath = "serv/image/Hangman_9_vies-removebg-preview.png"
	default:
		Data.ImagePath = "serv/image/Hangman_0_vies-removebg-preview.png"
	}
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

	guessedLetter := r.FormValue("LettreARecuperer")
	guessedLetter = formateLetter(guessedLetter)
	if guessedLetter == Data.Word {
		Victory(w, r)
	} else {
		if !letterAlreadyGuessed(guessedLetter) {
			Data.LettreUsed = append(Data.LettreUsed, guessedLetter)
			for i, char := range Data.Word {
				if string(char) == guessedLetter {

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
			image(Data.Try)
			nothere = true
			win = true
			for _, i := range Data.TabHidden {
				if i == "_" {
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
			Home(w, r)
		}
	}
}
