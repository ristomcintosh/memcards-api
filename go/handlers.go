package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var decks = []Deck{
	{Id: "1", Name: "World Capitals", Flashcards: flashcards1},
	{Id: "2", Name: "Basic Portuguese", Flashcards: flashcards2},
}

var flashcards1 = []Flashcard{
	{Id: "1", Front: "France", Back: "Paris", DeckId: "1"},
	{Id: "2", Front: "Japan", Back: "Tokyo", DeckId: "1"},
	{Id: "3", Front: "Italy", Back: "Rome", DeckId: "1"},
	{Id: "4", Front: "Brazil", Back: "Brasilia", DeckId: "1"},
	{Id: "5", Front: "Canada", Back: "Ottawa", DeckId: "1"},
}

var flashcards2 = []Flashcard{
	{Id: "6", Front: "Hello", Back: "Olá", DeckId: "2"},
	{Id: "7", Front: "Thank you", Back: "Obrigado", DeckId: "2"},
	{Id: "8", Front: "Yes", Back: "Sim", DeckId: "2"},
	{Id: "9", Front: "No", Back: "Não", DeckId: "2"},
	{Id: "10", Front: "Goodbye", Back: "Adeus", DeckId: "2"},
}

func GetDecks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(decks)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["deckId"]

	deck := findById(id, decks)

	if deck == nil {
		w.WriteHeader(http.StatusNotFound)
		message := fmt.Sprintf("Deck: %v not found", id)
		w.Write([]byte(message))
		return
	}

	err := json.NewEncoder(w).Encode(deck)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func findById(deckId string, decks []Deck) *Deck {
	for _, deck := range decks {
		if deck.Id == deckId {return &deck}
	}
	return nil
}