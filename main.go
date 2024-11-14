package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	Data.Lettre = "Bonjour"
	http.HandleFunc("/", Home)
	fmt.Println("(http://localhost:8080) - server started on port", port)
	http.ListenAndServe(port, nil)
	fs := http.FileServer(http.Dir("serv/"))
	http.Handle("serv/", http.StripPrefix("serv/", fs))
}
