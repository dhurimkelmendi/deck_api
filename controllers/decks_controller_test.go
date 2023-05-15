package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dhurimkelmendi/deck_api/controllers"
	"github.com/dhurimkelmendi/deck_api/fixtures"
	"github.com/go-chi/chi"
)

func TestDeckController(t *testing.T) {
	t.Parallel()
	fixture := fixtures.GetFixturesDefaultInstance()

	ctrl := controllers.GetControllersDefaultInstance()
	deck := fixture.Deck.CreateDeck(t)

	t.Run("create deck", func(t *testing.T) {
		t.Run("create partial deck", func(t *testing.T) {
			r := chi.NewRouter()
			r.Post("/api/v1/decks", ctrl.Decks.CreateDeck)

			bBuf := bytes.NewBuffer([]byte(fmt.Sprintf(`{"shuffled":false, "cards":["AH","AD","AS","AC"]}`)))
			req := httptest.NewRequest(http.MethodPost, "/api/v1/decks", bBuf)

			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			if res.Code != http.StatusCreated {
				t.Fatalf("expected http status code of 200 but got: %+v, %+v", res.Code, res.Body.String())
			}
		})
		t.Run("create full deck", func(t *testing.T) {
			r := chi.NewRouter()
			r.Post("/api/v1/decks", ctrl.Decks.CreateDeck)

			bBuf := bytes.NewBuffer([]byte("{}"))
			req := httptest.NewRequest(http.MethodPost, "/api/v1/decks", bBuf)

			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			if res.Code != http.StatusCreated {
				t.Fatalf("expected http status code of 200 but got: %+v, %+v", res.Code, res.Body.String())
			}
		})
	})
	t.Run("get deck", func(t *testing.T) {
		r := chi.NewRouter()
		URL := fmt.Sprintf("/api/v1/decks/%s", deck.ID.String())
		r.Get("/api/v1/decks/{id}", ctrl.Decks.GetDeckByID)

		req := httptest.NewRequest(http.MethodGet, URL, nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Fatalf("expected http status code of 200 but got: %+v, %+v", res.Code, res.Body.String())
		}
	})

	t.Run("draw card from deck", func(t *testing.T) {
		r := chi.NewRouter()
		URL := fmt.Sprintf("/api/v1/decks/%s", deck.ID.String())
		r.Patch("/api/v1/decks/{id}", ctrl.Decks.DrawCardFromDeck)

		bBuf := bytes.NewBuffer([]byte(""))
		req := httptest.NewRequest(http.MethodPatch, URL, bBuf)

		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Fatalf("expected http status code of 200 but got: %+v, %+v", res.Code, res.Body.String())
		}

		body := make(map[string]interface{})
		dec := json.NewDecoder(strings.NewReader(res.Body.String()))
		err := dec.Decode(&body)
		if err != nil {
			t.Fatalf("error decoding response body: %+v", err)
		}
	})
}
