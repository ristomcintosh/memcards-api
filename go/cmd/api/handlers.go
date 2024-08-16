package main

import (
	"encoding/json"
	"errors"
	"memcards-api/internal/data"
	"net/http"
	"strconv"

	// "github.com/google/uuid"
	"github.com/gorilla/mux"
)

type APIResponse struct {
	Message string
	Errors  []string
}

func (app *application) GetDecks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	decks, dbErr := app.db.GetAllDecks()

	if dbErr != nil {
		app.serverError(w, dbErr)
		return
	}

	err := json.NewEncoder(w).Encode(decks)

	if err != nil {
		app.serverError(w, err)
		return
	}

}

func (app *application) GetDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["deckId"]

	deck, dbErr := app.db.GetDeckByID(id)

	if dbErr != nil {
		if errors.Is(dbErr, data.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, dbErr)
		}
		return
	}

	err := json.NewEncoder(w).Encode(deck)

	if err != nil {
		app.serverError(w, err)
		return
	}
}

type CreateDeckInput struct {
	Name string `json:"name"`
}

func (app *application) CreateDeck(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name string `json:"name"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	// if req.Name == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("name is required"))
	// 	return
	// }

	newDeck, _ := app.db.CreateDeck(input.Name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newDeck)
	w.WriteHeader(http.StatusCreated)
}

func (app *application) UpdateDeck(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	if input.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("name is required"))
		return
	}

	vars := mux.Vars(r)
	// TODO handle error
	id, _ := strconv.Atoi(vars["deckId"])

	deck, _ := app.db.UpdateDeck(uint(id), input.Name)

	// if deck == nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	message := fmt.Sprintf("Deck: %v not found", id)
	// 	json.NewEncoder(w).Encode(APIResponse{
	// 		Message: message,
	// 	})
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deck)
}

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
