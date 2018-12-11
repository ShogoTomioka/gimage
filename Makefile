



NAME := ganalyze
VERSION := v0.1
GOVERSION := go1.11


SRCS    := $(shell find . -type f -name '*.go')
PKGS := $(shell go list ./...)

.PHONY: init
init:


.PHONY: build
build: test
	go build -a -v -o $(NAME) main.go

bin/$(NAME): $(SRCS)
	go build -a -v -o bin/$(NAME)

.PHONY: test
test: lint
	go test $(PKGS)

.PHONY: lint
lint:
	golint $(SRCS)

.PHONY: clean
clean:
	rm -rf ./bin
