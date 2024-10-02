# expense-tracker-api-go

A simple expense tracker in Go

- Clean architecture inspired from [https://github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- Dependency injection using `wire`

```text
├── cmd
│   ├── container.go # singleton container
│   ├── main.go  # application startup point   
│   ├── wire.go  # Injection
│   └── wire_gen.go  # Generate by Injection
├── config
│   └── config.go
├── domain          # Domain goes here
│   ├── expense.go 
│   └── user.go
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
└── service         # Business logics go here
    ├── expense_service.go
    └── user_service.go
```
