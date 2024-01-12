GOCMD = go
GOBUILD = $(GOCMD) build

BINARY_NAME = greenlight

run: build
	./$(BINARY_NAME)

build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/api

create-migration:
	migrate create -ext sql -dir migrations/ -seq $(NAME)
