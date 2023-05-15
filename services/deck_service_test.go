package services_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/dhurimkelmendi/deck_api/fixtures"
	"github.com/dhurimkelmendi/deck_api/payloads"
	"github.com/dhurimkelmendi/deck_api/services"
)

func TestDeckService(t *testing.T) {
	t.Parallel()
	fixture := fixtures.GetFixturesDefaultInstance()

	service := services.GetDeckServiceDefaultInstance()
	deck := fixture.Deck.CreateDeck(t)

	ctx := context.Background()

	t.Run("create deck", func(t *testing.T) {
		t.Run("create partial deck", func(t *testing.T) {
			deckToCreate := &payloads.CreateDeckPayload{}
			deckToCreate.Shuffled = gofakeit.Bool()

			cardCodes := make([]string, 0, 5)
			cardCodes = append(cardCodes, "2S")
			cardCodes = append(cardCodes, "2C")
			cardCodes = append(cardCodes, "2D")
			cardCodes = append(cardCodes, "2H")
			deckToCreate.CardCodes = cardCodes

			createedDeck, err := service.CreateDeck(ctx, deckToCreate)
			if err != nil {
				t.Fatalf("error while creating deck %+v", err)
			}
			deckToCreateModel := deckToCreate.ToDeckModel()
			deckToCreateModel.ID = createedDeck.ID
			if !deckToCreateModel.Equals(createedDeck) {
				t.Fatalf("create deck failed: %+v \n received: %+v, %+v", deckToCreateModel, createedDeck, err)
			}
		})
		t.Run("create full deck", func(t *testing.T) {
			deckToCreate := &payloads.CreateDeckPayload{}
			createdDeck, err := service.CreateDeck(ctx, deckToCreate)
			if err != nil {
				t.Fatalf("error while creating deck %+v", err)
			}
			if createdDeck.Shuffled != false {
				t.Fatalf("deck should not be shuffled when created")
			}
			if createdDeck.Remaining != 52 {
				t.Fatalf("deck should be full if no cards are specified")
			}
		})
	})
	t.Run("get deck by id", func(t *testing.T) {
		_, err := service.GetDeckByID(deck.ID)
		if err != nil {
			t.Fatalf("could not retreive existing deck by ID: %d, %+v", deck.ID, err)
		}
	})

	t.Run("draw card from deck", func(t *testing.T) {
		drawnCard, err := service.DrawCardFromDeck(ctx, deck.ID)
		if err != nil {
			t.Fatalf("update deck failed: %+v", err)
		}
		updatedDeck, err := service.GetDeckByID(deck.ID)
		if err != nil {
			t.Fatalf("could not retreive existing deck by ID: %d, %+v", deck.ID, err)
		}
		if updatedDeck.Remaining != deck.Remaining-1 {
			t.Fatalf("expected card to be drawn from deck, cards in deck not updated")
		}
		cardsRemaining := payloads.MapCardCodesToCards(updatedDeck.Cards)
		for _, c := range cardsRemaining {
			if c.Code == drawnCard.Code {
				t.Fatalf("expected card to be drawn from deck, card is still in deck")
			}
		}
		if drawnCard.Code == "" {
			t.Fatalf("expected card to be drawn from deck, got empty card")
		}
	})

}
