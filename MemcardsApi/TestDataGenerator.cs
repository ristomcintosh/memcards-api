using MemcardsApi.Models;

namespace MemcardsApi;

public static class TestDataGenerator
{

  public static List<Deck> Generate()
  {
    var decks = new List<Deck>();
    var deck1Id = "deck-1";
    var deck1 = new Deck(deck1Id, "Capitals",
[
    new Flashcard(Guid.NewGuid().ToString(), "Portugal", "Lisbon", deck1Id),
    new Flashcard(Guid.NewGuid().ToString(), "Spain", "Madrid", deck1Id),
    new Flashcard(Guid.NewGuid().ToString(), "France", "Paris", deck1Id),
    new Flashcard(Guid.NewGuid().ToString(), "Germany", "Berlin", deck1Id),
    new Flashcard(Guid.NewGuid().ToString(), "Italy", "Rome", deck1Id)
]);

    var deck2Id = "deck-2";
    var deck2 = new Deck(deck2Id, "Portuguese",
[
    new Flashcard(Guid.NewGuid().ToString(), "Hello", "Olá", deck2Id),
    new Flashcard(Guid.NewGuid().ToString(), "Goodbye", "Adeus", deck2Id),
    new Flashcard(Guid.NewGuid().ToString(), "Please", "Por favor", deck2Id),
    new Flashcard(Guid.NewGuid().ToString(), "Thank you", "Obrigado", deck2Id),
    new Flashcard(Guid.NewGuid().ToString(), "Sorry", "Desculpa", deck2Id)
]);

    decks.Add(deck1);
    decks.Add(deck2);

    return decks;
  }

}
