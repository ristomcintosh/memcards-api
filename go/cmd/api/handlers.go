package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"memcards-api/internal/data"
	"net/http"
	"strconv"
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

func (app *application) CreateFlashcard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deckId, _ := strconv.Atoi(vars["deckId"])

	var input struct {
		Front string `json:"front"`
		Back  string `json:"back"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	newFlashcard, _ := app.db.CreateFlashcard(uint(deckId), input.Front, input.Back)

	// if deck == nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	message := fmt.Sprintf("Deck: %v not found", deckId)
	// 	json.NewEncoder(w).Encode(APIResponse{
	// 		Message: message,
	// 	})
	// 	return
	// }

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newFlashcard)
}
