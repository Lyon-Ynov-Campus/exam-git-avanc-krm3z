package main

import (
	"fmt"
	"log"
	"net/http"
)

<<<<<<< HEAD
=======
func URLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>URL visitée</h1><p>Vous avez visité : %s</p>", r.URL.Path)
}

>>>>>>> 4ed5570 (Initialize feature-url)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Site Web en Go</h1><p>Page d'accueil</p>")
	})
<<<<<<< HEAD
=======
	
	http.HandleFunc("/url", URLHandler)
>>>>>>> 4ed5570 (Initialize feature-url)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
