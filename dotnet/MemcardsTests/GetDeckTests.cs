namespace MemcardsTests;

using System.Net.Http.Json;
using MemcardsApi.Models;
using FluentAssertions;
using Microsoft.AspNetCore.Http;
using System.Net;
using Microsoft.AspNetCore.Mvc.Testing;

public class GetDeckTests : IClassFixture<WebApplicationFactory<Program>>
{
    private readonly WebApplicationFactory<Program> _factory;

    public GetDeckTests(WebApplicationFactory<Program> factory)
    {
        _factory = factory;

    }

    [Fact]
    public async Task TestGetDecks()
    {
        var response = await _factory.CreateClient().GetAsync("/decks");
        var responseData = await response.Content.ReadFromJsonAsync<List<Deck>>();
        Assert.True(response.IsSuccessStatusCode);
        Assert.Equal("Capitals", responseData?[0].Name);
    }

    [Fact]
    public async Task TestGetDeck()
    {
        var testDeck = MemcardsApi.TestDataGenerator.Generate()[0];
        var response = await _factory.CreateClient().GetAsync("/decks/deck-1");
        var responseData = await response.Content.ReadFromJsonAsync<Deck>();
        Assert.True(response.IsSuccessStatusCode);
        Assert.Equal("Capitals", responseData?.Name);
        if (responseData != null)
        {
            testDeck.Should().BeEquivalentTo(responseData, options => options
            .Excluding(deck => deck.Id)
            .Excluding(deck => deck.Path.EndsWith("Id")));
        }
    }

    [Fact]
    public async Task TestGetDeckNotFound()
    {
        var response = await _factory.CreateClient().GetAsync("/decks/bad-id");

        Assert.Equal(HttpStatusCode.NotFound, response.StatusCode);
    }
}