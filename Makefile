test:
	go test -v ./...

lint:
	go vet ./... && golint ./...

build:
	go build && gouml i && mv ./file.puml ./documents/classdiagram.puml

