package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", whatUpHandler)
	r.HandleFunc("/decks", GetDecks)

	fmt.Println("severing on port 5757")
	log.Fatal(http.ListenAndServe(":5757", r))
}

func whatUpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("What up"))
}