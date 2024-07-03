using MemcardsApi.Models;

namespace MemcardsApi.Services;

public interface IDbService
{
  Deck? GetDeck(string deckId);
  Deck? UpdateDeck(string deckId, DeckName deckName);
  List<Deck> GetDecks();
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
}
