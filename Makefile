build:
	@go build -o bin/ecom.exe cmd/main.go

test:
	@go test -v ./...
	
run: build
	@.\bin\ecom.exe