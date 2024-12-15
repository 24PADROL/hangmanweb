package main

import (
	h "hangmanweb/hangmanweb"
	"net/http"
)

func main() {
	h.Init()
	h.Web()

	http.HandleFunc("/", h.Menu)     //page menu
	http.HandleFunc("/home", h.Home) //page game
	http.HandleFunc("/victory", h.Victory)
	http.HandleFunc("/lose", h.Lose)

	http.ListenAndServe(":8080", nil)

}
