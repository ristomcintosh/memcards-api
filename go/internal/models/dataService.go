package models

import "gorm.io/gorm"

type DataService struct {
	*gorm.DB
}

func (ds *DataService) GetAllDecks() ([]Deck, error) {
	var decks []Deck
	result := ds.Model(&Deck{}).Preload("Flashcards").Find(&decks)
	return decks, result.Error
}
