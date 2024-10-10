build:
	@ go build -trimpath -o app .

go-generate:
	go generate ./...

swag:
	swag init -g http_server.go -d cmd,internal/http,domain
	swag fmt

test:
	go test ./... --cover -timeout 30s