package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type application struct {
	db gorm.DB
}

func main() {
	app := application{
		db: *dbSetup(),
	}

	handlers := app.routes()

	fmt.Println("listening on port 5757")
	log.Fatal(http.ListenAndServe(":5757", handlers))
}

var _decks = []Deck{
	{Name: "World Capitals"},
	// {Name: "Basic Portuguese"},
}

// var _flashcards1 = []Flashcard{
// 	{Id: "1", Front: "France", Back: "Paris", DeckID: "1"},
// 	{Id: "2", Front: "Japan", Back: "Tokyo", DeckID: "1"},
// 	{Id: "3", Front: "Italy", Back: "Rome", DeckID: "1"},
// 	{Id: "4", Front: "Brazil", Back: "Brasilia", DeckID: "1"},
// 	{Id: "5", Front: "Canada", Back: "Ottawa", DeckID: "1"},
// }

// var _flashcards2 = []Flashcard{
// 	{Id: "6", Front: "Hello", Back: "Olá", DeckID: "2"},
// 	{Id: "7", Front: "Thank you", Back: "Obrigado", DeckID: "2"},
// 	{Id: "8", Front: "Yes", Back: "Sim", DeckID: "2"},
// 	{Id: "9", Front: "No", Back: "Não", DeckID: "2"},
// 	{Id: "10", Front: "Goodbye", Back: "Adeus", DeckID: "2"},
// }

func dbSetup() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Exec("DROP TABLE decks")
	db.AutoMigrate(Deck{}, Flashcard{})

	db.Create(&_decks)

	worldCapitals := _decks[0]
	worldCapitalsCards := []Flashcard{
		{Front: "France", Back: "Paris", DeckID: worldCapitals.ID},
		{Front: "Japan", Back: "Tokyo", DeckID: worldCapitals.ID},
		{Front: "Italy", Back: "Rome", DeckID: worldCapitals.ID},
		{Front: "Brazil", Back: "Brasilia", DeckID: worldCapitals.ID},
		{Front: "Canada", Back: "Ottawa", DeckID: worldCapitals.ID},
	}
	db.Create(worldCapitalsCards)

	return db
}
