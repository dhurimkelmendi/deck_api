package models

import (
	"net/http"
	"reflect"

	uuid "github.com/satori/go.uuid"
)

// Deck is a struct that represents a db row of the Decks table
type Deck struct {
	tableName struct{}  `pg:"decks"`
	ID        uuid.UUID `pg:"id,pk,type:uuid"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
	Cards     []string  `json:"cards"`
}

// Equals compares two instances of type Deck
func (d *Deck) Equals(secondDeck *Deck) bool {
	if d.ID != secondDeck.ID {
		return false
	}
	if d.Shuffled != secondDeck.Shuffled {
		return false
	}
	if d.Remaining != secondDeck.Remaining {
		return false
	}
	return reflect.DeepEqual(d.Cards, secondDeck.Cards)
}

// Render is used by go-chi/renderer
func (d *Deck) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
