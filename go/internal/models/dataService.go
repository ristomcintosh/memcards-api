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

func (ds *DataService) GetDeckByID(id string) (*Deck, error) {
	var deck *Deck
	result := ds.Model(&deck).Find(&deck, id)

	if result.RowsAffected == 0 {
		deck = nil
	}

	return deck, result.Error
}
