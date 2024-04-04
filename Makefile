GOCMD = go
GOBUILD = $(GOCMD) build

BINARY_NAME = greenlight

confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

build:
	$(GOBUILD) -o $(BINARY_NAME) -ldflags='-s' ./cmd/api

run/api: build
	./$(BINARY_NAME)

db/psql:
	psql ${GREENLIGHT_DB_DSN}

db/migrations/new:
	migrate create -ext sql -dir migrations/ -seq ${name}

# migrate -path ./migrations -database postgres://postgres:postgres@localhost/greenlight?sslmode=disable up
db/migrations/up: confirm
	@echo 'Running migrations...'
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up

audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...
