build:
	@ go build -trimpath -o app ./cmd/

go-generate:
	go generate ./...

test:
	go test ./...