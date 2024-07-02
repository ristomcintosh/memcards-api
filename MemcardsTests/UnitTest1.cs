namespace MemcardsTests;

using System.Net.Http.Json;
using MemcardsApi.Models;
using Microsoft.AspNetCore.Mvc.Testing;
using FluentAssertions;
using Microsoft.AspNetCore.Http;
using System.Net;

public class HelloWorldTests : IClassFixture<WebApplicationFactory<Program>>
{
    private readonly HttpClient client;

    public HelloWorldTests(WebApplicationFactory<Program> factory)
    {
        client = factory.CreateClient();
    }

    [Fact]
    public async Task TestRootEndpoint()
    {
        var response = await client.GetStringAsync("/");

        Assert.Equal("Hello World!", response);
    }

    [Fact]
    public async Task TestGetDecks()
    {
        var response = await client.GetAsync("/decks");
        var responseData = await response.Content.ReadFromJsonAsync<List<Deck>>();
        Assert.True(response.IsSuccessStatusCode);
        Assert.Equal("Capitals", responseData?[0].Name);
    }

    [Fact]
    public async Task TestGetDeck()
    {
        var testDeck = MemcardsApi.TestDataGenerator.Generate()[0];
        var response = await client.GetAsync("/decks/deck-1");
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
        var response = await client.GetAsync("/decks/bad-id");

        Assert.Equal(HttpStatusCode.NotFound, response.StatusCode);
    }

}