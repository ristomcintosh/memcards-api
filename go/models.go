package main

import "gorm.io/gorm"

type Identifiable interface {
	GetId() uint
}

type Repository[T Identifiable] struct {
	items []T
}

func (r *Repository[T]) FindById(id uint) *T {
	for i, item := range r.items {
		if item.GetId() == id {
			return &r.items[i]
		}
	}
	return nil
}

func (r *Repository[T]) Create(newItem T) {
	r.items = append(r.items, newItem)
}

type Deck struct {
	gorm.Model
	Name       string `json:"name"`
	Flashcards []Flashcard
}

func (d Deck) GetId() uint {
	return d.ID
}

type Flashcard struct {
	gorm.Model
	Front  string `json:"front"`
	Back   string `json:"back"`
	DeckID uint   `json:"deckId"`
}

func (f Flashcard) GetId() uint {
	return f.ID
}
