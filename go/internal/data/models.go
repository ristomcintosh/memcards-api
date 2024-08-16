package data

import "gorm.io/gorm"

type Flashcard struct {
	gorm.Model
	Front  string `json:"front"`
	Back   string `json:"back"`
	DeckID uint   `json:"deckId"`
}

type Deck struct {
	gorm.Model
	ID         uint   `gorm:"primarykey"`
	Name       string `json:"name"`
	Flashcards []Flashcard
}
