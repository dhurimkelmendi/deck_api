# Deck API
## Usage instructions
- Add necessary ENV variables(check `/config/config.go` to see what is necessary)
- Create DB with the `config.DatabaseName` name
- Run `make db_upgrade` to run migrations
- Feel free to run the API using `make serve` or run the test suite using `make test`

## Available commands
### `make serve`
Runs the server

### `make db_upgrade`
Runs the pending migrations

### `make db_reset`
Resets db to initial state

### `make db_downgrade`
Rolls back the last migration that was ran in the db

### `make test`
Runs test suite
