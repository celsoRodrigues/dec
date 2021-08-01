#include .env

#go section
mod:
	go mod init github.com/celsoRodrigues/dec
tidy:
	go mod tidy

test:
	go test -v --bench=. --benchmem

vendor:
	go mod vendor -v

build:
	go build -o ./bin/kubectl-dec ./cmd
	@echo "binary file in ./bin directory"

run:
	go run ./cmd/*.go

install: tidy build
	cp ./bin/kubectl-dec /usr/local/bin

all: mod tidy vendor build install run

uninstall:
	rm -rf /usr/local/bin/kubectl-dec

clean: uninstall
	rm ./bin/kubectl-dec
	rm go.mod go.sum
	rm ./vendor -rf
