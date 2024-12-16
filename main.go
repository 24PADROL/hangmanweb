package main

import (
	h "hangmanweb/hangmanweb"
	"net/http"
    "log"
)

func main() {
	h.Init()
	h.Web()

	http.HandleFunc("/", h.Menu)     //page menu
	http.HandleFunc("/home", h.Home) //page game
	http.HandleFunc("/victory", h.Victory)
	http.HandleFunc("/lose", h.Lose)

	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("Erreur lors du démarrage du serveur HTTPS :", err)
	}

}
