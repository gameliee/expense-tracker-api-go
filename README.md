# expense-tracker-api-go

A simple expense tracker in Go

- Clean architecture inspired from [https://github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- Dependency injection using `wire`

```text
├── main.go          # main
├── cmd
│   ├── container.go # singleton container
│   ├── database.go    
│   ├── http_server.go
│   └── app.go       # init application
├── config
│   └── config.go
├── domain          # Domain goes here
│   ├── expense.go 
│   └── user.go
├── docs            # Swagger docs
├── go.mod
├── go.sum
├── internal        # Infrastructure
│   ├── http        # Controller/Delivery
│   │   ├── expense_handler.go
│   │   └── user_handler.go
│   ├── repository  # Repository
│   │   └── sqlite
│   │       ├── expense_repository.go
│   │       └── user_repository.go
│   └── workers
├── tests           # Test all the layers
├── service         # Business logics go here
│   ├── expense_service.go
│   └── user_service.go
└── tools
    ├── dependency_injector.go   # Field DI
    └── dependency_injector_test.go
```

Dependencies

```bash
brew install mockery
```
