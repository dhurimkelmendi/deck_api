package services

import (
	"context"

	"github.com/dhurimkelmendi/deck_api/db"
	"github.com/dhurimkelmendi/deck_api/models"
	"github.com/dhurimkelmendi/deck_api/payloads"

	"github.com/go-pg/pg/v10"
	uuid "github.com/satori/go.uuid"
)

// DeckService is a struct that contains references to the db and the StatelessAuthenticationProvider
type DeckService struct {
	db *pg.DB
}

var deckServiceDefaultInstance *DeckService

// GetDeckServiceDefaultInstance returns the default instance of DeckService
func GetDeckServiceDefaultInstance() *DeckService {
	if deckServiceDefaultInstance == nil {
		deckServiceDefaultInstance = &DeckService{
			db: db.GetDefaultInstance().GetDB(),
		}
	}

	return deckServiceDefaultInstance
}

// GetDeckByID returns the requested deck by id
func (s *DeckService) GetDeckByID(deckID uuid.UUID) (*models.Deck, error) {
	return s.getDeckByID(deckID)
}
func (s *DeckService) getDeckByID(deckID uuid.UUID) (*models.Deck, error) {
	deck := &models.Deck{}
	switch err := s.db.Model(deck).Where("id = ?", deckID).Select(); err {
	case pg.ErrNoRows:
		return deck, db.ErrNoMatch
	default:
		return deck, err
	}
}

// CreateDeck creates a deck using the provided payload
func (s *DeckService) CreateDeck(ctx context.Context, createDeck *payloads.CreateDeckPayload) (*models.Deck, error) {
	deck := &models.Deck{}
	if err := createDeck.Validate(); err != nil {
		return deck, err
	}

	var err error
	err = s.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		deck, err = s.createDeck(tx, createDeck)
		return err
	})
	if err != nil {
		return deck, err
	}
	return deck, err
}

func (s *DeckService) createDeck(dbSession *pg.Tx, createDeck *payloads.CreateDeckPayload) (*models.Deck, error) {
	deck := createDeck.ToDeckModel()
	deck.ID = uuid.NewV4()
	if len(deck.Cards) == 0 {
		deck.Cards = payloads.GetAllCardCodes()
		deck.Shuffled = false
		deck.Remaining = len(deck.Cards)
	}

	_, err := dbSession.Model(deck).Insert()
	if err != nil {
		return deck, err
	}

	return deck, nil
}

// DrawCardFromDeck updates the deck by id using the provided payload
func (s *DeckService) DrawCardFromDeck(ctx context.Context, deckID uuid.UUID) (*payloads.Card, error) {
	var card *payloads.Card
	var err error
	s.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		card, err = s.drawCardFromDeck(tx, deckID)
		return err
	})

	return card, err
}
func (s *DeckService) drawCardFromDeck(dbSession *pg.Tx, deckID uuid.UUID) (*payloads.Card, error) {
	deck, err := s.GetDeckByID(deckID)
	if err != nil {
		return &payloads.Card{}, db.ErrNoMatch
	}

	if len(deck.Cards) == 0 {
		return &payloads.Card{}, nil
	}
	cardCode := deck.Cards[0]
	cardCodesToMap := make([]string, 0, 1)
	cardCodesToMap = append(cardCodesToMap, cardCode)
	cards := payloads.MapCardCodesToCards(cardCodesToMap)
	card := cards[0]

	deck.Cards = deck.Cards[1:]
	deck.Remaining--

	if _, err := dbSession.Model(deck).Where("id = ?", deck.ID).Update(); err != nil {
		if err == pg.ErrNoRows {
			return card, db.ErrNoMatch
		}
		return card, err
	}
	return card, nil
}
