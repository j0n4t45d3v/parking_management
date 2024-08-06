all: build 

BIN_PATH="./bin/api"
MAIN_GO="./cmd/api/main.go"

bin-dir:
	@if [ ! -d "./bin" ]; then mkdir -p $(BIN_PATH); fi

build: bin-dir
	go build -o $(BIN_PATH) $(MAIN_GO)
