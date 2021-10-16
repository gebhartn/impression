.SILENT:

HOST=localhost
PORT=8080
CONTAINER_PORT=8080

export APIURL=http://$(HOST):$(PORT)/api

GOOS=linux
GOARCH=amd64
APP=impress
APP_STATIC=$(APP)-static
LDFLAGS="-w -s -extldflags=-static"

download:
	go mod download

generate:
	go generate .

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -race -o $(APP) .

build-static:
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -ldflags=$(LDFLAGS) -o $(APP_STATIC) .

build-image:
	docker build -t $(APP) .

run:
	go run -race .

run-container:
	chmod o+w ./database && docker run -p 8080:8080 -v "$(pwd)"/database:/app/database $(APP):latest

test:
	go test -v ./...

test-race:
	go test -v -race ./...
