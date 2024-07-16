package main

type Deck struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Flashcards []Flashcard
}

type Flashcard struct {
	Id string `json:"id"`
	Front string `json:"front"`
	Back string `json:"back"`
	DeckId string `json:"deckId"`
}