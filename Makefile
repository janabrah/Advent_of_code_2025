BIN_DIR := bin
# Default command name if you don't pass CMD=...
CMD ?= playground

.PHONY: run fmt vet clean

# Main target: format, vet, build selected command, then run it
run: fmt vet $(BIN_DIR)/$(CMD)
	@$(BIN_DIR)/$(CMD)

fmt:
	go fmt ./...

vet:
	go vet ./...

# Build rule: builds bin/<name> from cmd/<name>/*.go
$(BIN_DIR)/%: cmd/%/*.go
	mkdir -p $(BIN_DIR)
	go build -o $@ ./cmd/$*

clean:
	rm -rf $(BIN_DIR)

.PHONY: run-% build-%

run-%: fmt vet $(BIN_DIR)/%
	@$(BIN_DIR)/$*

build-%: $(BIN_DIR)/%
	@echo "Built $@"