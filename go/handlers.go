package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

var deckRepository = Repository[Deck]{decks}

type APIResponse struct {
	Message string
	Errors  []string
}

func GetDecks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(deckRepository.items)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["deckId"]

	deck := deckRepository.FindById(id)

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

type CreateDeckInput struct {
	Name string `json:"name"`
}

func CreateDeck(w http.ResponseWriter, r *http.Request) {
	var req CreateDeckInput

	json.NewDecoder(r.Body).Decode(&req)

	if req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("name is required"))
		return
	}

	newDeck := Deck{
		Id:         uuid.New().String(),
		Name:       req.Name,
		Flashcards: []Flashcard{},
	}

	deckRepository.Create(newDeck)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newDeck)
	w.WriteHeader(http.StatusCreated)
}

type UpdateDeckInput struct {
	Name string `json:"name"`
}

func UpdateDeck(w http.ResponseWriter, r *http.Request) {
	var req UpdateDeckInput

	json.NewDecoder(r.Body).Decode(&req)

	if req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("name is required"))
		return
	}

	vars := mux.Vars(r)
	id := vars["deckId"]

	deck := deckRepository.FindById(id)

	if deck == nil {
		w.WriteHeader(http.StatusNotFound)
		message := fmt.Sprintf("Deck: %v not found", id)
		json.NewEncoder(w).Encode(APIResponse{
			Message: message,
		})
		return
	}

	deck.Name = req.Name

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deck)
}

type createFlashcardInput struct {
	Front string `json:"front" validate:"required"`
	Back  string `json:"back" validate:"required"`
}

func CreateFlashcard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deckId := vars["deckId"]

	var reqBody createFlashcardInput
	json.NewDecoder(r.Body).Decode(&reqBody)

	validationErrors, err := ValidateRequestBody(reqBody)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(validationErrors) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Message: "Invalid data provided",
			Errors:  validationErrors,
		})
		return
	}

	var deck = deckRepository.FindById(deckId)

	if deck == nil {
		w.WriteHeader(http.StatusNotFound)
		message := fmt.Sprintf("Deck: %v not found", deckId)
		json.NewEncoder(w).Encode(APIResponse{
			Message: message,
		})
		return
	}

	newFlashcard := Flashcard{
		Id:     uuid.New().String(),
		Front:  reqBody.Front,
		Back:   reqBody.Back,
		DeckId: deckId,
	}

	deck.Flashcards = append(deck.Flashcards, newFlashcard)

	json.NewEncoder(w).Encode(newFlashcard)
	w.WriteHeader(http.StatusCreated)
}
