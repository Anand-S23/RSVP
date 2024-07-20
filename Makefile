build:
	@go build -o bin/rsvp backend/main.go

run: build
	@./bin/rsvp

test:
	@go test -v ./...
