package fixtures

import "github.com/dhurimkelmendi/deck_api/db"

// Fixtures is a struct that contains references to all fixture instances.
type Fixtures struct {
	Deck *DeckFixture
}

var fixturesDefaultInstance *Fixtures

// GetFixturesDefaultInstance returns the default instance of Fixtures
func GetFixturesDefaultInstance() *Fixtures {
	if fixturesDefaultInstance == nil {
		// Purposeful pre-initialize
		_ = db.GetDefaultInstance()

		fixturesDefaultInstance = &Fixtures{
			Deck: GetDeckFixtureDefaultInstance(),
		}
	}
	return fixturesDefaultInstance
}
