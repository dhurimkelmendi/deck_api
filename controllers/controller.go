package controllers

import "github.com/dhurimkelmendi/deck_api/api"

// Controllers is a struct that contains references to all controller instances.
type Controllers struct {
	Decks *DecksController
}

// Controller is a struct that contains references to error components and responders
type Controller struct {
	errCmp    api.ErrorComponentFn
	responder *api.Responder
}

var controllersDefaultInstance *Controllers

// GetControllersDefaultInstance returns default instances of all available Controllers
func GetControllersDefaultInstance() *Controllers {
	if controllersDefaultInstance == nil {
		controllersDefaultInstance = &Controllers{
			Decks: GetDecksControllerDefaultInstance(),
		}
	}
	return controllersDefaultInstance
}
