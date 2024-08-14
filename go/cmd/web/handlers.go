package main

import (
	"encoding/json"
	// "fmt"
	"net/http"
	// "github.com/google/uuid"
	// "github.com/gorilla/mux"
)

type APIResponse struct {
	Message string
	Errors  []string
}

func (app *application) GetDecks() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		decks, dbErr := app.db.GetAllDecks()

		if dbErr != nil {
			app.errorLog.Print("something went with the db!")
			return
		}

		err := json.NewEncoder(w).Encode(decks)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

// func GetDeck(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["deckId"]

// 	deck := deckRepository.FindById(id)

// 	if deck == nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		message := fmt.Sprintf("Deck: %v not found", id)
// 		w.Write([]byte(message))
// 		return
// 	}

// 	err := json.NewEncoder(w).Encode(deck)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }

// type CreateDeckInput struct {
// 	Name string `json:"name"`
// }

// func CreateDeck(w http.ResponseWriter, r *http.Request) {
// 	var req CreateDeckInput

// 	json.NewDecoder(r.Body).Decode(&req)

// 	if req.Name == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("name is required"))
// 		return
// 	}

// 	newDeck := Deck{
// 		Id:         uuid.New().String(),
// 		Name:       req.Name,
// 		Flashcards: []Flashcard{},
// 	}

// 	deckRepository.Create(newDeck)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(newDeck)
// 	w.WriteHeader(http.StatusCreated)
// }

// type UpdateDeckInput struct {
// 	Name string `json:"name"`
// }

// func UpdateDeck(w http.ResponseWriter, r *http.Request) {
// 	var req UpdateDeckInput

// 	json.NewDecoder(r.Body).Decode(&req)

// 	if req.Name == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("name is required"))
// 		return
// 	}

// 	vars := mux.Vars(r)
// 	id := vars["deckId"]

// 	deck := deckRepository.FindById(id)

// 	if deck == nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		message := fmt.Sprintf("Deck: %v not found", id)
// 		json.NewEncoder(w).Encode(APIResponse{
// 			Message: message,
// 		})
// 		return
// 	}

// 	deck.Name = req.Name

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(deck)
// }

// type createFlashcardInput struct {
// 	Front string `json:"front" validate:"required"`
// 	Back  string `json:"back" validate:"required"`
// }

// func CreateFlashcard(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	deckId := vars["deckId"]

// 	var reqBody createFlashcardInput
// 	json.NewDecoder(r.Body).Decode(&reqBody)

// 	validationErrors, err := ValidateRequestBody(reqBody)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	if len(validationErrors) != 0 {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(APIResponse{
// 			Message: "Invalid data provided",
// 			Errors:  validationErrors,
// 		})
// 		return
// 	}

// 	var deck = deckRepository.FindById(deckId)

// 	if deck == nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		message := fmt.Sprintf("Deck: %v not found", deckId)
// 		json.NewEncoder(w).Encode(APIResponse{
// 			Message: message,
// 		})
// 		return
// 	}

// 	newFlashcard := Flashcard{
// 		Id:     uuid.New().String(),
// 		Front:  reqBody.Front,
// 		Back:   reqBody.Back,
// 		DeckID: deckId,
// 	}

// 	deck.Flashcards = append(deck.Flashcards, newFlashcard)

// 	json.NewEncoder(w).Encode(newFlashcard)
// 	w.WriteHeader(http.StatusCreated)
// }
