.DEFAULT_GOAL := help

.PHONY: all
# generate all
all: protoc gen build

# ==============================================================================
# Includes

include scripts/make-rules/all.mk

# ==============================================================================
# Goals

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: tidy
# tidy
tidy: 
	@go mod tidy

.PHONY: protoc
# Generate api proto files.
protoc:
	$(MAKE) gen.protoc

.PHONY: gen
# generate
gen: tidy
	go generate ./...

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
