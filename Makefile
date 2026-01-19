all: test build
	cp ./bin/pingenemy 

build:
	go build -ldflags "-w" -o bin/pingenemy

run:
	go run main.go

clean-test-cache:
	go clean -testcache

test: clean-test-cache
	go list -f '{{.Dir}} {{.ImportPath}}' ./... | awk 'system("ls " $$1 "/*_test.go > /dev/null 2>&1") == 0 {print $$2}' | xargs go test

execute: build
	./bin/pingenemy