package payloads

import (
	"net/http"

	"github.com/dhurimkelmendi/deck_api/models"
	uuid "github.com/satori/go.uuid"
)

// Card represents a card object
type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// MapCardCodesToCards convert a card code slice to a card struct slice
func MapCardCodesToCards(cardCodes []string) []*Card {
	availableCards := getListOfAvailableCards()
	cards := make([]*Card, 0, len(cardCodes))
	for _, cardCode := range cardCodes {
		for _, card := range availableCards {
			if card.Code == cardCode {
				cards = append(cards, card)
			}
		}
	}
	return cards
}

// GetAllCardCodes convert a card code slice to a card struct slice
func GetAllCardCodes() []string {
	availableCards := getListOfAvailableCards()
	cards := make([]string, 0, len(availableCards))
	for _, card := range availableCards {
		cards = append(cards, card.Code)
	}
	return cards
}

// Render is used by go-chi/renderer
func (c *Card) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// DeckDetails simple response object
type DeckDetails struct {
	ID        uuid.UUID `json:"id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
	Cards     []*Card   `json:"cards"`
}

// Render is used by go-chi/renderer
func (u *DeckDetails) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// MapDeckToDeckDetails convert a deck model to a payload response
func MapDeckToDeckDetails(deck *models.Deck) *DeckDetails {
	return &DeckDetails{
		ID:        deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
		Cards:     MapCardCodesToCards(deck.Cards),
	}
}

// CreateDeckPayload for creating a new deck
type CreateDeckPayload struct {
	Shuffled  bool     `json:"shuffled"`
	CardCodes []string `json:"cards"`
}

// ToDeckModel converts an instance of type *RegisterDeckPayload to *models.Deck type
func (d *CreateDeckPayload) ToDeckModel() *models.Deck {
	return &models.Deck{
		Remaining: len(d.CardCodes),
		Shuffled:  d.Shuffled,
		Cards:     d.CardCodes,
	}
}

// Validate ensures that all the required fields are present in an instance of *RegisterDeckPayload
func (d *CreateDeckPayload) Validate() error {
	return nil
}

// Render is used by go-chi/renderer
func (d *CreateDeckPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func getListOfAvailableCards() []*Card {
	return []*Card{
		{
			Suit:  "hearts",
			Value: "2",
			Code:  "2H",
		},
		{
			Suit:  "hearts",
			Value: "3",
			Code:  "3H",
		},
		{
			Suit:  "hearts",
			Value: "4",
			Code:  "4H",
		},
		{
			Suit:  "hearts",
			Value: "5",
			Code:  "5H",
		},
		{
			Suit:  "hearts",
			Value: "6",
			Code:  "6H",
		},
		{
			Suit:  "hearts",
			Value: "7",
			Code:  "7H",
		},
		{
			Suit:  "hearts",
			Value: "8",
			Code:  "8H",
		},
		{
			Suit:  "hearts",
			Value: "9",
			Code:  "9H",
		},
		{
			Suit:  "hearts",
			Value: "10",
			Code:  "10H",
		},
		{
			Suit:  "hearts",
			Value: "J",
			Code:  "JH",
		},
		{
			Suit:  "hearts",
			Value: "Q",
			Code:  "QH",
		},
		{
			Suit:  "hearts",
			Value: "K",
			Code:  "KH",
		},
		{
			Suit:  "hearts",
			Value: "A",
			Code:  "AH",
		},
		{
			Suit:  "diamonds",
			Value: "2",
			Code:  "2D",
		},
		{
			Suit:  "diamonds",
			Value: "3",
			Code:  "3D",
		},
		{
			Suit:  "diamonds",
			Value: "4",
			Code:  "4D",
		},
		{
			Suit:  "diamonds",
			Value: "5",
			Code:  "5D",
		},
		{
			Suit:  "diamonds",
			Value: "6",
			Code:  "6D",
		},
		{
			Suit:  "diamonds",
			Value: "7",
			Code:  "7D",
		},
		{
			Suit:  "diamonds",
			Value: "8",
			Code:  "8D",
		},
		{
			Suit:  "diamonds",
			Value: "9",
			Code:  "9D",
		},
		{
			Suit:  "diamonds",
			Value: "10",
			Code:  "10D",
		},
		{
			Suit:  "diamonds",
			Value: "J",
			Code:  "JD",
		},
		{
			Suit:  "diamonds",
			Value: "Q",
			Code:  "QD",
		},
		{
			Suit:  "diamonds",
			Value: "K",
			Code:  "KD",
		},
		{
			Suit:  "diamonds",
			Value: "A",
			Code:  "AD",
		},
		{
			Suit:  "clubs",
			Value: "2",
			Code:  "2C",
		},
		{
			Suit:  "clubs",
			Value: "3",
			Code:  "3C",
		},
		{
			Suit:  "clubs",
			Value: "4",
			Code:  "4C",
		},
		{
			Suit:  "clubs",
			Value: "5",
			Code:  "5C",
		},
		{
			Suit:  "clubs",
			Value: "6",
			Code:  "6C",
		},
		{
			Suit:  "clubs",
			Value: "7",
			Code:  "7C",
		},
		{
			Suit:  "clubs",
			Value: "8",
			Code:  "8C",
		},
		{
			Suit:  "clubs",
			Value: "9",
			Code:  "9C",
		},
		{
			Suit:  "clubs",
			Value: "10",
			Code:  "10C",
		},
		{
			Suit:  "clubs",
			Value: "J",
			Code:  "JC",
		},
		{
			Suit:  "clubs",
			Value: "Q",
			Code:  "QC",
		},
		{
			Suit:  "clubs",
			Value: "K",
			Code:  "KC",
		},
		{
			Suit:  "clubs",
			Value: "A",
			Code:  "AC",
		},
		{
			Suit:  "spades",
			Value: "2",
			Code:  "2S",
		},
		{
			Suit:  "spades",
			Value: "3",
			Code:  "3S",
		},
		{
			Suit:  "spades",
			Value: "4",
			Code:  "4S",
		},
		{
			Suit:  "spades",
			Value: "5",
			Code:  "5S",
		},
		{
			Suit:  "spades",
			Value: "6",
			Code:  "6S",
		},
		{
			Suit:  "spades",
			Value: "7",
			Code:  "7S",
		},
		{
			Suit:  "spades",
			Value: "8",
			Code:  "8S",
		},
		{
			Suit:  "spades",
			Value: "9",
			Code:  "9S",
		},
		{
			Suit:  "spades",
			Value: "10",
			Code:  "10S",
		},
		{
			Suit:  "spades",
			Value: "J",
			Code:  "JS",
		},
		{
			Suit:  "spades",
			Value: "Q",
			Code:  "QS",
		},
		{
			Suit:  "spades",
			Value: "K",
			Code:  "KS",
		},
		{
			Suit:  "spades",
			Value: "A",
			Code:  "AS",
		},
	}
}
