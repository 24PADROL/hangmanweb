package hangmanweb

import (
	h "hangmanweb/hangmanweb"
)

const port = ":8080"

func main() {

	randomWord()
	http.HandleFunc("/", Home)
	http.HandleFunc("/input", Input)
	fmt.Println("(http://localhost:8080) - server started on port", port)
	http.ListenAndServe(port, nil)
	fs := http.FileServer(http.Dir("serv/"))
	http.Handle("serv/", http.StripPrefix("serv/", fs))
}
