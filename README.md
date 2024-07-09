# Memcards API

This repository contains **multiple implementations** of the same API using different programming languages and frameworks. It serves as a playground for learning and experimenting with new technologies.

## API Endpoints

All implementations follow the same API specification for [Memcards (a simple flashcard app)](https://github.com/ristomcintosh/memcards-v2). The endpoints are as follows:

### Decks

- `GET /decks` - Retrieve a list of all decks
- `GET /decks/{deckId}` - Retrieve a specific deck
- `POST /decks` - Create a new deck
- `PUT /decks/{deckId}` - Update a deck
- `DELETE /decks/{deckId}` - Delete a deck

### Flashcards

- `POST /decks/{deckId}/flashcards` - Add a new flashcard to deck
- `PUT /decks/{deckId}/flashcards/{flashcardId}` - Update a flashcard
- `DELETE /decks/{deckId}/flashcards/{flashcardId}` - Delete a flashcard
