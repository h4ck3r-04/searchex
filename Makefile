build:
	go build -o searchex main.go

run:
	go run main.go

test:
	go test ./...

clean:
	rm -rf searchex
