package main

import (
	"fmt"
	"net/http"
)

var port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)
	fmt.Println("(http://localhost:8080) - server started on port", port)
	http.ListenAndServe(port, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Me parlez pas")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Me parlez pas")
}
