


VERSION := v0.1
GOVERSION := go1.11
NAME := gimage

TYPE := A
SRCS := $(shell find ./lib -type f -name '*.go')

.PHONY: run
run: build
	./${NAME} ${TYPE}

.PHONY: build
build: lint
	go build -o gimage main.go

.PHONY: lint
lint:
	golint $(SRCS)

.PHONY: clean
clean:
	rm -rf ./testdata/filtered.png
	rm -rf ./testdata/binary.png
	rm -r  ./${NAME}
