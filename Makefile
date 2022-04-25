build:
	go fmt
	go build

clean:
	go clean

rel:
	go fmt
	go build -ldflags="-s -w"
