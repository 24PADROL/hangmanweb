package hangmanweb

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type DataForm struct {
	LettreUsed []string // Lettres déjà utilisées
	Words      []string // Liste de mots chargée depuis le fichier
	NameFill   string   // Nom du fichier de mots
	Word       string   // Mot sélectionné
	TabHidden  []string // Mot caché (sous forme "_ _ _ _")
	Letter     string   // Lettre soumise par l'utilisateur
}

var Data DataForm
var win bool = false

// Nom du fichier par défaut contenant les mots
var nameFill string = "motsimple.txt"

// Port par défaut du serveur
const port = ":8080"

// Génère un mot aléatoire depuis le fichier
func randomWord() {
	fichier, err := os.Open(nameFill)
	if err != nil {
		fmt.Println("Erreur d'ouverture du fichier:", err)
		return
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		Data.Words = append(Data.Words, scanner.Text())
	}

	if len(Data.Words) == 0 {
		fmt.Println("Le fichier est vide.")
		return
	}

	// Sélectionner un mot aléatoire
	rand.Seed(time.Now().UnixNano()) // Initialiser le générateur de nombres aléatoires
	Data.Word = strings.ToLower(Data.Words[rand.Intn(len(Data.Words))])
	fmt.Println("Mot sélectionné :", Data.Word) // Debugging
}

// Initialise le mot caché sous forme "_ _ _ _"
func printHidden() {
	Data.TabHidden = []string{}
	for _, char := range Data.Word {
		if char == ' ' {
			Data.TabHidden = append(Data.TabHidden, " ") // Gérer les espaces
		} else {
			Data.TabHidden = append(Data.TabHidden, "_")
		}
	}
}

// Initialise les données pour une nouvelle partie
func Init() {
	Data = DataForm{} // Réinitialiser les données
	randomWord()
	printHidden()
}

// Point d'entrée principal du serveur web
func Web() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/input", Input)
	http.HandleFunc("/victory", Victory)
	http.HandleFunc("/reset", Reset) // Nouveau gestionnaire pour réinitialiser
	fs := http.FileServer(http.Dir("serv/"))
	http.Handle("/serv/", http.StripPrefix("/serv/", fs))

	fmt.Println("(http://localhost:8080) - Le serveur a démarré sur le port", port)
	http.ListenAndServe(port, nil)
}


// Gestionnaire pour la page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		html := `
		<!DOCTYPE html>
		<html lang="fr">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Jeu du Pendu</title>
			<style>
				body {
					background-color: #f0f8ff;
					font-family: Arial, sans-serif;
					color: #333;
					display: flex;
					justify-content: center;
					align-items: center;
					flex-direction: column;
					height: 100vh;
					margin: 0;
				}
				h1 {
					color: #007acc;
				}
				form {
					margin: 20px 0;
				}
				input[type="text"] {
					padding: 10px;
					font-size: 16px;
					border: 1px solid #007acc;
					border-radius: 5px;
					width: 200px;
					text-align: center;
				}
				input[type="text"]:focus {
					outline: none;
					border-color: #005a9e;
					box-shadow: 0 0 5px rgba(0, 90, 158, 0.5);
				}
				.feedback {
					margin-top: 20px;
					font-size: 18px;
				}
			</style>
		</head>
		<body>
			<h1>Jeu du Pendu</h1>
			<div class="feedback">
				<p>Mot caché : ` + strings.Join(Data.TabHidden, " ") + `</p>
				<p>Lettres utilisées : ` + strings.Join(Data.LettreUsed, ", ") + `</p>
			</div>
			<form method="POST" action="/input" id="letterForm">
				<input name="letter" type="text" placeholder="Entrez une lettre" maxlength="1" required id="letterInput">
			</form>
			<script>
				// Focus automatique sur le champ de texte
				window.onload = () => {
					document.getElementById('letterInput').focus();
				};

				// Soumission automatique après saisie d'une lettre
				const inputField = document.getElementById('letterInput');
				const form = document.getElementById('letterForm');

				inputField.addEventListener('input', () => {
					if (inputField.value.length === 1) {
						form.submit();
					}
				});
			</script>
		</body>
		</html> `
		w.Write([]byte(html))
	}
}

// Gestionnaire pour les entrées utilisateur
func Input(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		letter := strings.ToLower(r.FormValue("letter"))

		// Validation de la lettre
		if len(letter) != 1 || letter < "a" || letter > "z" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Vérifier si la lettre a déjà été utilisée
		for _, used := range Data.LettreUsed {
			if used == letter {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}

		Data.LettreUsed = append(Data.LettreUsed, letter)

		// Mettre à jour le mot caché
		found := false
		for i, char := range Data.Word {
			if string(char) == letter {
				Data.TabHidden[i] = letter
				found = true
			}
		}

		// Vérifier si le joueur a gagné
		if strings.Join(Data.TabHidden, "") == Data.Word {
			win = true
			http.Redirect(w, r, "/victory", http.StatusSeeOther)
			return
		}

		if !found {
			fmt.Println("Lettre incorrecte :", letter) // Debugging
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// Gestionnaire pour la page de victoire
func Victory(w http.ResponseWriter, r *http.Request) {
	if win {
		html := `
		<!DOCTYPE html>
		<html lang="fr">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Victoire</title>
			<style>
				body {
					background-color: #f0f8ff;
					font-family: Arial, sans-serif;
					color: #333;
					text-align: center;
					padding: 20px;
				}
				h1 {
					font-size: 3rem;
					color: #28a745;
					margin: 20px 0;
					animation: pop 1s ease-in-out;
				}
				p {
					font-size: 1.5rem;
					margin: 20px 0;
				}
				a {
					display: inline-block;
					margin-top: 20px;
					padding: 10px 20px;
					background-color: #007bff;
					color: white;
					text-decoration: none;
					font-size: 1.2rem;
					border-radius: 5px;
					transition: background-color 0.3s ease;
				}
				a:hover {
					background-color: #0056b3;
				}
				@keyframes pop {
					0% {
						transform: scale(0.8);
						opacity: 0;
					}
					100% {
						transform: scale(1);
						opacity: 1;
					}
				}
			</style>
		</head>
		<body>
			<h1>🎉 Félicitations ! 🎉</h1>
			<p>Vous avez deviné le mot : <strong>` + Data.Word + `</strong></p>
			<a href="/reset">Rejouer</a>
		</body>
		</html>
		`
		w.Write([]byte(html))
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func Reset(w http.ResponseWriter, r *http.Request) {
	Init() // Réinitialise les données de jeu
	http.Redirect(w, r, "/", http.StatusSeeOther) // Redirige vers la page d'accueil
}

