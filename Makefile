test:
	go test -v ./...

lint:
	go vet ./... && golint ./...

build:
	go build 

