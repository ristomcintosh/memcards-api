using MemcardsApi.Models;

namespace MemcardsApi.Services;

public interface IDbService
{
  Deck? GetDeck(string deckId);
  Deck? UpdateDeck(string deckId, DeckName deckName);
  List<Deck> GetDecks();

  int DeleteDeck(string deckId);

  Flashcard? CreateFlashcard(string deckId, Flashcard flashcard);

  Flashcard? UpdateFlashcard(string deckId, string flashcardId, Flashcard flashcard);

  int DeleteFlashcard(string deckId, string flashcardId);
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

  public Flashcard? UpdateFlashcard(string deckId, string flashcardId, Flashcard flashcard)
  {
    var deck = decks.FirstOrDefault((deck) => deck.Id == deckId);
    if (deck == null) return null;
    var existingFlashcard = deck.Flashcards.FirstOrDefault((flashcard) => flashcard.Id == flashcardId);
    if (existingFlashcard == null) return null;
    existingFlashcard.Front = flashcard.Front;
    existingFlashcard.Back = flashcard.Back;
    return existingFlashcard;
  }

  public int DeleteFlashcard(string deckId, string flashcardId)
  {
    var deck = decks.FirstOrDefault((deck) => deck.Id == deckId);
    if (deck == null) return 0;
    var flashcard = deck.Flashcards.FirstOrDefault((flashcard) => flashcard.Id == flashcardId);
    if (flashcard == null) return 0;
    deck.Flashcards.Remove(flashcard);
    return 1;
  }
}
