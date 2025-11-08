GO := go
BIN_NAME := Pollen
CMD_DIR := ./cmd/pollen


build:
	@$(GO) build -o $(BIN_NAME) $(CMD_DIR)

run: build
	@./$(BIN_NAME)