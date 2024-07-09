

using MemcardsApi.Services;
using MemcardsApi.Models;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddSingleton<IDbService, InMemoryDB>();
var app = builder.Build();



app.MapGet("/decks", (IDbService service) =>
{
  return TypedResults.Ok(service.GetDecks());
});

app.MapGet("decks/{deckId}", (string deckId, IDbService service) =>
{
  var deck = service.GetDeck(deckId);
  if (deck == null)
  {
    return Results.NotFound();
  }
  return TypedResults.Ok(deck);
});

app.MapPut("/decks/{deckId}", (string deckId, DeckName deckName, IDbService service) =>
{
  var existingDeck = service.UpdateDeck(deckId, deckName);
  if (existingDeck == null)
  {
    return Results.NotFound();
  }
  return TypedResults.Ok(existingDeck);
});

app.MapDelete("/decks/{deckId}", (string deckId, IDbService service) =>
{
  var deck = service.DeleteDeck(deckId);
  if (deck.Equals(0))
  {
    return Results.NotFound();
  }
  return Results.NoContent();
});

app.MapPost("/decks/{deckId}/flashcards", (string deckId, Flashcard flashcard, IDbService service) =>
{
  var createdFlashcard = service.CreateFlashcard(deckId, flashcard);
  if (createdFlashcard == null)
  {
    return Results.NotFound();
  }

  return Results.Created("", createdFlashcard);

});

app.MapPut("/decks/{deckId}/flashcards/{flashcardId}", (string deckId, string flashcardId, Flashcard flashcard, IDbService service) =>
{
  var updatedFlashcard = service.UpdateFlashcard(deckId, flashcardId, flashcard);
  if (updatedFlashcard == null)
  {
    return Results.NotFound();
  }
  return TypedResults.Ok(updatedFlashcard);
});

app.MapDelete("/decks/{deckId}/flashcards/{flashcardId}", (string deckId, string flashcardId, IDbService service) =>
{
  var flashcard = service.DeleteFlashcard(deckId, flashcardId);
  if (flashcard.Equals(0))
  {
    return Results.NotFound();
  }
  return Results.NoContent();
});

app.Run();


public partial class Program { }