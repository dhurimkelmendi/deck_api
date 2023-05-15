package migrations

import (
	"github.com/go-pg/migrations/v8"
	"github.com/sirupsen/logrus"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		logrus.Infoln("Creating decks table")
		_, err := db.Exec(`
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

		CREATE TABLE decks (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			shuffled boolean DEFAULT FALSE,
			remaining int,
			cards jsonb
		);`)
		return err
	}, func(db migrations.DB) error {
		logrus.Infoln("Dropping decks table")
		_, err := db.Exec(`
			DROP TABLE IF EXISTS decks CASCADE;
			DROP EXTENSION "uuid-ossp";
		`)
		return err
	})
}
