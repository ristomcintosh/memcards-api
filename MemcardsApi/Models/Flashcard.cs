namespace MemcardsApi.Models;

public class Flashcard(string Id, string Front, string Back, string DeckId)
{
  public string Id { get; set; } = Id;
  public string Front { get; set; } = Front;
  public string Back { get; set; } = Back;
  public string DeckId { get; set; } = DeckId;
}