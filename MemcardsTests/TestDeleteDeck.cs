using System.Net;
using Microsoft.AspNetCore.Mvc.Testing;

namespace MemcardsTests;

public class TestDeleteDeck : IClassFixture<WebApplicationFactory<Program>>
{
  private readonly WebApplicationFactory<Program> _factory;

  public TestDeleteDeck(WebApplicationFactory<Program> factory)
  {
    _factory = factory;

  }

  [Fact]
  public async Task Test_Delete_Deck()
  {
    var response = await _factory.CreateClient().DeleteAsync("/decks/deck-1");

    Assert.Equal(HttpStatusCode.NoContent, response.StatusCode);
  }

}
