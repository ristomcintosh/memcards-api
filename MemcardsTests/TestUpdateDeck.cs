namespace MemcardsTests;
using System.Net.Http.Json;
using MemcardsApi.Models;
using FluentAssertions;
using Microsoft.AspNetCore.Http;
using System.Net;
using Microsoft.AspNetCore.Mvc.Testing;

public class TestUpdateDeck : IClassFixture<WebApplicationFactory<Program>>
{
  private readonly WebApplicationFactory<Program> _factory;

  public TestUpdateDeck(WebApplicationFactory<Program> factory)
  {
    _factory = factory;

  }

  [Fact]
  public async Task Test_Update_Deck()
  {
    var response = await _factory.CreateClient().PutAsJsonAsync("/decks/deck-1", new { Name = "Updated Name" });

    var responseData = await response.Content.ReadFromJsonAsync<Deck>();
    Assert.Equal(HttpStatusCode.OK, response.StatusCode);
    responseData.Should().NotBeNull();
    Assert.Equal("Updated Name", responseData?.Name);
  }

  [Fact]
  public async Task Test_Update_Deck_Not_Found()
  {
    var response = await _factory.CreateClient().GetAsync("/decks/bad-id");

    Assert.Equal(HttpStatusCode.NotFound, response.StatusCode);
  }

}
