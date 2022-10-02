all: test lint build

test:
	go test -v ./...

lint:
	go vet ./... && golint ./...

GOPATH=$$(go env GOPATH)
GOPLANTUML="$(GOPATH)/bin/goplantuml"
build:
	go build && $(GOPLANTUML) -show-aggregations -aggregate-private-members -title BulitienBoard -recursive -output ./documents/calssdiagram.puml .