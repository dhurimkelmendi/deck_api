package fixtures

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/dhurimkelmendi/deck_api/db"
	"github.com/dhurimkelmendi/deck_api/models"
	"github.com/dhurimkelmendi/deck_api/payloads"
	"github.com/dhurimkelmendi/deck_api/services"
	"github.com/go-pg/pg/v10"
)

// DeckFixture is a struct that contains references to the db and DeckService
type DeckFixture struct {
	db          *pg.DB
	deckService *services.DeckService
}

var deckFixtureDefaultInstance *DeckFixture

// GetDeckFixtureDefaultInstance returns the default instance of DeckFixture
func GetDeckFixtureDefaultInstance() *DeckFixture {
	if deckFixtureDefaultInstance == nil {
		deckFixtureDefaultInstance = &DeckFixture{
			db:          db.GetDefaultInstance().GetDB(),
			deckService: services.GetDeckServiceDefaultInstance(),
		}
	}

	if deckFixtureDefaultInstance.deckService == nil {
		deckFixtureDefaultInstance.deckService = services.GetDeckServiceDefaultInstance()
	}

	return deckFixtureDefaultInstance
}

// CreateDeck creates a deck with fake data
func (f *DeckFixture) CreateDeck(t *testing.T) *models.Deck {
	deck := &payloads.CreateDeckPayload{}
	deck.Shuffled = gofakeit.Bool()

	cardCodes := make([]string, 0, 5)
	cardCodes = append(cardCodes, "AS")
	cardCodes = append(cardCodes, "AC")
	cardCodes = append(cardCodes, "AD")
	cardCodes = append(cardCodes, "AH")
	deck.CardCodes = cardCodes

	ctx := context.Background()

	if f.deckService == nil {
		t.Log("CreateBuyerDeck: fixture.DeckService is nil!")
	}

	createdDeck, err := f.deckService.CreateDeck(ctx, deck)
	if err != nil {
		return nil
	}
	return createdDeck
}
