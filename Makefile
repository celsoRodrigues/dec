#include .env

#go section
tidy:
	go mod tidy

test:
	go test -v --bench=. --benchmem

vendor:
	go mod vendor -v

build:
	go build -o ./bin/kubectl-enc ./cmd
	@echo "binary file in ./bin directory"

run:
	go run ./cmd/*.go

install: build
	cp ./bin/kubectl-enc /usr/local/bin

all: tidy vendor build install run

uninstall:
	rm /usr/local/bin/kubectl-enc

clean: uninstall
	rm ./bin/kubectl-enc
	rm go.mod go.sum
	rm ./vendor -rf
