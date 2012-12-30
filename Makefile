all:
	go build -o bin/thingol src/*.go

run:
	go run src/*.go
