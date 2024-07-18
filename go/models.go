package main

type Identifiable interface {
	GetId() string
}

type Repository[T Identifiable] struct {
	items []T
}

func (r *Repository[T]) FindById(id string) *T {
	for i, item := range r.items {
		if item.GetId() == id	 {return &r.items[i]}
	}
	return nil
}

func (r *Repository[T]) Create(newItem T) {
	r.items = append(r.items, newItem)
}

type Deck struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Flashcards []Flashcard
}

func (d Deck) GetId() string {
	return d.Id
}

type Flashcard struct {
	Id string `json:"id"`
	Front string `json:"front"`
	Back string `json:"back"`
	DeckId string `json:"deckId"`
}

func (f Flashcard) GetId() string {
	return f.Id
}