package main

import (
	h "hangmanweb/hangmanweb"
)

func main() {
	h.Init()
	h.Web()

	// go func() {
	// 	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
	// 	})))
	// }()

	// err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	// if err != nil {
	// 	log.Fatal("Erreur lors du démarrage du serveur HTTPS :", err)
	// }
}
