.PHONY: run build test clean

run:
	go run ./cmd/pricepulse

build:
	go build -o bin/pricepulse ./cmd/pricepulse

test:
	go test ./...

clean:
	rm -rf bin