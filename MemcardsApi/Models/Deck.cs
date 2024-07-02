namespace MemcardsApi.Models;

public record Deck(string Id, string Name, List<Flashcard> Flashcards);
