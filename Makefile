all:
	go build -ldflags "-w" -o bin/pingenemy

run:
	go run main.go

execute: all
	./bin/pingenemy