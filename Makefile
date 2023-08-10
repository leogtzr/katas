default: build ;

test:
	go test ./...

clean:
	@echo "Cleaning up build junk"
	go clean

build:
	go build

cover:
	go test -v ./... -coverprofile ./cp.out
	go tool cover -html=cp.out
