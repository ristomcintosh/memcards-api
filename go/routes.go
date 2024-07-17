package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/decks", GetDecks).Methods(http.MethodGet)
	r.HandleFunc("/decks/{deckId}", GetDeck).Methods(http.MethodGet)
	r.HandleFunc("/decks", CreateDeck).Methods(http.MethodPost)

	return r
}