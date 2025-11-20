all:
	go build -ldflags "-w" -o bin/pingenemy

run:
	go run

execute:
	./pingenemy