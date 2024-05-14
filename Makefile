build:
	go mod download
	go mod verify
	CGO_ENABLED=0 go build -o phaidra-assessment main.go

test:
	go test -v ./...

run: build
	./phaidra-assessment
