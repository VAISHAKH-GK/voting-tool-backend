build:
	go build -o bin/voting-tool cmd/main.go

run: build
	./bin/voting-tool
