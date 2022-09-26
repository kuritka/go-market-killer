SHELL := bash

ifndef NO_COLOR
YELLOW=\033[0;33m
CYAN=\033[1;36m
RED=\033[31m
# no color
NC=\033[0m
endif

GOLIC_VERSION  ?= v0.7.2
GOLINT_VERSION  ?= v1.45.0


test:
	go test ./... --cover

.PHONY: lint
lint:
	goimports -w ./
	@echo -e "\n$(YELLOW)Running the linters$(NC)"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINT_VERSION)
	$(GOBIN)/golangci-lint run

# updates source code with license headers
.PHONY: license
license:
	@echo -e "\n$(YELLOW)Injecting the license$(NC)"
	@go install github.com/AbsaOSS/golic@$(GOLIC_VERSION)
	$(GOBIN)/golic inject -t apache2
#	$(GOBIN)/golic remove -t mit

check: lint test

all: check
