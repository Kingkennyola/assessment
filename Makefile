build:
	go mod download
	go mod verify
	CGO_ENABLED=0 go build -o assessment main.go

test:
	go test -v ./...

run: build
	./phaidra-assessment

build-docker:
	docker build -t ghcr.io/kingkennyola/assessment/assessment:0.0.1 .
