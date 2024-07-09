using MemcardsApi.Models;

namespace MemcardsApi.Services;

public interface IDbService
{
  Deck? GetDeck(string deckId);
  Deck? UpdateDeck(string deckId, DeckName deckName);
  List<Deck> GetDecks();

  int DeleteDeck(string deckId);

  Flashcard? CreateFlashcard(string deckId, Flashcard flashcard);
}


public class InMemoryDB : IDbService
{
  private List<Deck> decks = [];
  public InMemoryDB()
  {
    decks = TestDataGenerator.Generate();
  }
  public Deck? GetDeck(string deckId)
  {
    return decks.FirstOrDefault((deck) => deck.Id == deckId);
  }

  public List<Deck> GetDecks()
  {
    return decks;
  }

  public Deck? UpdateDeck(string deckId, DeckName deckName)
  {
    var existingDeck = decks.FirstOrDefault((deck) => deck.Id == deckId);
    if (existingDeck == null)
    {
      return null;
    }
    existingDeck.Name = deckName.Name;
    return existingDeck;
  }

  public int DeleteDeck(string deckId)
  {
    var deck = decks.FirstOrDefault((deck) => deck.Id == deckId);
    if (deck != null)
    {
      decks.Remove(deck);
      return 1;
    };
    return 0;
  }

  public Flashcard? CreateFlashcard(string deckId, Flashcard flashcard)
  {
    var deck = decks.FirstOrDefault((deck) => deck.Id == deckId);

    if (deck == null) return null;
    var newFlashcard = new Flashcard(Guid.NewGuid().ToString(), flashcard.Front, flashcard.Back, deckId);
    deck.Flashcards.Add(newFlashcard);
    return newFlashcard;
  }
}
