BIN=crtchkr
GO=$(shell command -v go)

all: build

build:
	$(GO) build -o $(BIN) .

run:
	./$(BIN)

clean:
	rm -f $(BIN)