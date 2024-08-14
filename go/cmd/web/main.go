package main

import (
	"log"
	"memcards-api/internal/models"
	"net/http"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type application struct {
	db       *models.DataService
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := dbSetup()

	if err != nil {
		errorLog.Fatal(err)
	}

	app := application{
		db:       db,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	handlers := app.routes()

	infoLog.Println("listening on port 5757")
	log.Fatal(http.ListenAndServe(":5757", handlers))
}

var decks = []models.Deck{
	{Name: "World Capitals"},
	{Name: "Basic Portuguese"},
}

func dbSetup() (*models.DataService, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Exec("DROP TABLE decks")
	db.Exec("DROP TABLE flashcards")
	db.AutoMigrate(models.Deck{}, models.Flashcard{})

	db.Create(&decks)

	worldCapitals := decks[0]
	worldCapitalsCards := []models.Flashcard{
		{Front: "France", Back: "Paris", DeckID: worldCapitals.ID},
		{Front: "Japan", Back: "Tokyo", DeckID: worldCapitals.ID},
		{Front: "Italy", Back: "Rome", DeckID: worldCapitals.ID},
		{Front: "Brazil", Back: "Brasilia", DeckID: worldCapitals.ID},
		{Front: "Canada", Back: "Ottawa", DeckID: worldCapitals.ID},
	}
	db.Create(worldCapitalsCards)

	portugueseBasic := decks[1]
	portugueseBasicCards := []models.Flashcard{
		{Front: "Hello", Back: "Olá", DeckID: portugueseBasic.ID},
		{Front: "Thank you", Back: "Obrigado", DeckID: portugueseBasic.ID},
		{Front: "Yes", Back: "Sim", DeckID: portugueseBasic.ID},
		{Front: "No", Back: "Não", DeckID: portugueseBasic.ID},
		{Front: "Goodbye", Back: "Adeus", DeckID: portugueseBasic.ID},
	}

	db.Create(portugueseBasicCards)

	return &models.DataService{DB: db}, nil
}
