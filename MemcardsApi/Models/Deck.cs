namespace MemcardsApi.Models;

public class Deck(string id, string name, List<Flashcard>? flashcards = null)
{
  public string Id { get; set; } = id;
  public string Name { get; set; } = name;
  public List<Flashcard> Flashcards { get; set; } = flashcards ?? [];
}


public class DeckName(string name)
{
  public string Name { get; set; } = name;
}