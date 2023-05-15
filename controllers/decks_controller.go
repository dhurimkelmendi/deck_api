package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dhurimkelmendi/deck_api/api"
	"github.com/dhurimkelmendi/deck_api/db"
	"github.com/dhurimkelmendi/deck_api/payloads"
	"github.com/dhurimkelmendi/deck_api/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"
)

// A DecksController handles HTTP requests that deal with deck.
type DecksController struct {
	Controller
	deckService *services.DeckService
}

var decksControllerDefaultInstance *DecksController

// GetDecksControllerDefaultInstance returns the default instance of DeckController.
func GetDecksControllerDefaultInstance() *DecksController {
	if decksControllerDefaultInstance == nil {
		decksControllerDefaultInstance = NewDeckController(services.GetDeckServiceDefaultInstance())
	}

	return decksControllerDefaultInstance
}

// NewDeckController create a new instance of a deck controller using the supplied deck service
func NewDeckController(deckService *services.DeckService) *DecksController {
	controller := Controller{
		errCmp:    api.NewErrorComponent(api.CmpController),
		responder: api.GetResponderDefaultInstance(),
	}
	return &DecksController{
		Controller:  controller,
		deckService: deckService,
	}
}

// CreateDeck creates a new deck and returns deck details with an authentication token
func (c *DecksController) CreateDeck(w http.ResponseWriter, r *http.Request) {
	errCtx := c.errCmp(api.CtxCreateDeck, r.Header.Get("X-Request-Id"))
	deck := &payloads.CreateDeckPayload{}
	if err := json.NewDecoder(r.Body).Decode(deck); err != nil {
		c.responder.Error(w, errCtx(api.ErrCreatePayload, errors.New("cannot decode deck payload")), http.StatusBadRequest)
		return
	}

	if err := deck.Validate(); err != nil {
		c.responder.Error(w, errCtx(api.ErrInvalidRequestPayload, errors.New("request body not valid, missing required fields")), http.StatusBadRequest)
		return
	}

	createdDeck, err := c.deckService.CreateDeck(context.Background(), deck)
	if err != nil {
		c.responder.Error(w, errCtx(api.ErrCreateDeck, err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	c.responder.JSON(w, r, createdDeck, http.StatusCreated)
}

// GetDeckByID returns the requested deck by id
func (c *DecksController) GetDeckByID(w http.ResponseWriter, r *http.Request) {
	errCtx := c.errCmp(api.CtxGetDeck, r.Header.Get("X-Request-Id"))
	urlDeckID := chi.URLParam(r, "id")
	deckID, err := uuid.FromString(urlDeckID)
	if err != nil {
		c.responder.Error(w, errCtx(api.ErrInvalidRequestParameter, fmt.Errorf("invalid deckId, %v", err)), http.StatusBadRequest)
		return
	}

	deck, err := c.deckService.GetDeckByID(deckID)
	if err != nil {
		if err == db.ErrNoMatch {
			c.responder.Error(w, errCtx(api.ErrDeckNotFound, errors.New("no deck with that id")), http.StatusNotFound)
		} else {
			c.responder.Error(w, errCtx(api.ErrGetDeck, err), http.StatusBadRequest)
		}
		return
	}

	var res render.Renderer
	res = payloads.MapDeckToDeckDetails(deck)

	if err := render.Render(w, r, res); err != nil {
		c.responder.Error(w, errCtx(api.ErrCreatePayload, errors.New("cannot serialize result")), http.StatusBadRequest)
		return
	}
}

// DrawCardFromDeck updates the current decks profile
func (c *DecksController) DrawCardFromDeck(w http.ResponseWriter, r *http.Request) {
	errCtx := c.errCmp(api.CtxDrawCardFromDeck, r.Header.Get("X-Request-Id"))
	urlDeckID := chi.URLParam(r, "id")
	deckID, err := uuid.FromString(urlDeckID)
	if err != nil {
		c.responder.Error(w, errCtx(api.ErrInvalidRequestParameter, fmt.Errorf("invalid deckId, %v", err)), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	defer r.Body.Close()

	drawnCard, err := c.deckService.DrawCardFromDeck(ctx, deckID)
	if err != nil {
		c.responder.Error(w, errCtx(api.ErrDrawCardFromDeck, err), http.StatusBadRequest)
		return
	}

	if err := render.Render(w, r, drawnCard); err != nil {
		c.responder.Error(w, errCtx(api.ErrDrawCardFromDeck, err), http.StatusBadRequest)
		return
	}
}
