

using MemcardsApi.Models;

var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", () => "Hello World!");

var decks = MemcardsApi.TestDataGenerator.Generate();

app.MapGet("/decks", () =>
{
  return TypedResults.Ok(decks);
});

app.MapGet("decks/{deckId}", (string deckId) =>
{
  var deck = decks.FirstOrDefault((deck) => deck.Id == deckId);
  if (deck == null)
  {
    return Results.NotFound();
  }
  return TypedResults.Ok(deck);
});

app.MapPut("/decks/{deckId}", (string deckId, DeckName deckName) =>
{
  var existingDeck = decks.FirstOrDefault((deck) => deck.Id == deckId);
  if (existingDeck == null)
  {
    return Results.NotFound();
  }
  existingDeck.Name = deckName.Name;
  return TypedResults.Ok(existingDeck);
});

app.Run();


public partial class Program { }