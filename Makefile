build:
	@ go build -trimpath -o app ./cmd/

go-generate:
	go generate ./...

swag:
	swag init -g http_server.go -d cmd,internal/http,domain
	swag fmt

test:
	go test ./... --cover