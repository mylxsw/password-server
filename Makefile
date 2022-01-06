LDFLAGS := -s -w -X main.Version=$(shell date "+%Y%m%d%H%M") -X main.GitRev=$(shell git rev-parse HEAD)

run:
	go run main.go

build:
	go build -race -ldflags "$(LDFLAGS)" -o bin/password-server main.go


.PHONY: run build
