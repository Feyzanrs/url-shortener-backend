// cmd/main.go

package main

import (
	"log"
	"net/http"
	"url-shortener/shortener" // Bu satırı kontrol edin
)

func main() {
	http.HandleFunc("/create", shortener.HandleCreate)
	http.HandleFunc("/r/", shortener.HandleRedirect)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
