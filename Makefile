build:
	@ go build -trimpath -o app ./cmd/

go-generate:
	go generate ./...

swag:
	swag init --dir cmd
	swag fmt

test:
	go test ./...