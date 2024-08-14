package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/decks", app.GetDecks()).Methods(http.MethodGet)
	r.HandleFunc("/decks/{deckId}", app.GetDeck).Methods(http.MethodGet)
	// r.HandleFunc("/decks", CreateDeck).Methods(http.MethodPost)
	// r.HandleFunc("/decks/{deckId}", UpdateDeck).Methods(http.MethodPut)
	// r.HandleFunc("/decks/{deckId}/flashcards", CreateFlashcard).Methods(http.MethodPost)

	return r
}
