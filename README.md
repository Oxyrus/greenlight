# Greenlight

Code from the book Let's Go Further!

## Setup

You must add your PostgreSQL connection to an environment variable:
```
export GREENLIGHT_DB="postgres://greenlight:greenlight@localhost/greenlight?sslmodle=disable"
```

## Migrations

You need the [migrate](https://github.com/golang-migrate/migrate) CLI tool. You can install it by running `$ brew install golang-migrate`.

### Create a new migration

`$ migrate create -seq -ext=.sql -dir=./migrations my_migration_name`.

### Run migrations

`$ migrate -path=./migrations -database=$GREENLIGHT_DB` This is why it's important to add the database connection as an environment variable.
