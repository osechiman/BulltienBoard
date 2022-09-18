all: test lint build

test:
	go test -v ./...

lint:
	go vet ./... && golint ./...

GOPATH=$$(go env GOPATH)
GOUML="$(GOPATH)/bin/gouml"
build:
	go build && $(GOUML) i && mv ./file.puml ./documents/classdiagram.puml

