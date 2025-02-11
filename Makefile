SHELL := bash # we want bash behaviour in all shell invocations
 
# COLORS
BOLD := \033[1m
NORMAL := \033[0m
GREEN := \033[1;32m

.DEFAULT_GOAL := help
HELP_TARGET_DEPTH ?= \#
help: # show how to get started & what targets are available
	@printf "This is a list of all the make targets that you can run, e.g. $(BOLD)make test$(NORMAL)\n\n"
	@awk -F':+ |$(HELP_TARGET_DEPTH)' '/^[0-9a-zA-Z._%-]+:+.+$(HELP_TARGET_DEPTH).+$$/ { printf "$(GREEN)%-20s\033[0m %s\n", $$1, $$3 }' $(MAKEFILE_LIST)
	@echo

# ------------------------------------------------------------------------------
# testing
# ------------------------------------------------------------------------------
.PHONY: fmt
fmt: # fmt for all the go files.
	@test -z $(shell echo $(shell go fmt $(shell go list ./... | grep -v /vendor/)) | tr -d "[:space:]")

.PHONY: test
test: # run all tests
	go test -shuffle=on -count 1 -race ./...

# ------------------------------------------------------------------------------
# running 
# ------------------------------------------------------------------------------
.PHONY: run
run: # execution of the main of the http-server api
	go run cmd/http/main.go


# ------------------------------------------------------------------------------
# building
# ------------------------------------------------------------------------------
.PHONY: build
build: # build the binary and saves it in the bin folder.
	CGO_ENABLED=0 GOOS=linux go build -o bin/http-api cmd/http/main.go

.PHONY: image
image: # build the docker image with tag booking-insights.
	docker build -t booking-insights .
