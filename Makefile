GOCMD = go
GOBUILD = $(GOCMD) build

BINARY_NAME = greenlight

run: build
	./$(BINARY_NAME)

build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/api
