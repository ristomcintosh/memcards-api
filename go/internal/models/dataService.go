package models

import (
	"errors"

	"gorm.io/gorm"
)

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

	err := ds.First(&deck, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return deck, nil
}
