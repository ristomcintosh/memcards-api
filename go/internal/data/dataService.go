package data

import (
	"errors"

	"gorm.io/gorm"
)

type DataService interface {
	GetAllDecks() ([]Deck, error)
	GetDeckByID(id string) (*Deck, error)
	CreateDeck(name string) (*Deck, error)
}

type GormOrm struct {
	*gorm.DB
}

func (orm *GormOrm) GetAllDecks() ([]Deck, error) {
	var decks []Deck
	result := orm.Model(&Deck{}).Preload("Flashcards").Find(&decks)
	return decks, result.Error
}

func (orm *GormOrm) GetDeckByID(id string) (*Deck, error) {
	var deck Deck

	err := orm.First(&deck, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &deck, nil
}

func (orm *GormOrm) CreateDeck(name string) (*Deck, error) {
	deck := &Deck{Name: name}
	err := orm.DB.Create(&deck).Error

	if err != nil {
		return nil, err
	}

	return deck, nil
}
