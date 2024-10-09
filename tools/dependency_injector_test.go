package tools_test

import (
	"fmt"
	"gamelieelearn/expense-tracker-api-go/tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ServiceA Interface and Implementation
type ServiceA interface {
	DoWorkA() string
}

type MyServiceA struct {
	ServiceB ServiceB `inject:"*tools_test.MyServiceB"` // Inject ServiceB
}

func (s *MyServiceA) DoWorkA() string {
	return "Service A is working"
}

// ServiceB Interface and Implementation
type ServiceB interface {
	DoWorkB() string
}

type MyServiceB struct {
	ServiceA ServiceA `inject:"*tools_test.MyServiceA"` // Field Injection using struct tag
}

func (s *MyServiceB) DoWorkB() string {
	return fmt.Sprintf("Service B is working, and %s", s.ServiceA.DoWorkA())
}

func TestDI(t *testing.T) {
	// Create a new container
	container := tools.NewContainer()

	// Register instances
	container.RegisterInstance(&MyServiceA{})
	container.RegisterInstance(&MyServiceB{})

	err := container.Build()
	assert.NoError(t, err)

	retrievedServiceA := container.Get((*MyServiceA)(nil)).(*MyServiceA)
	retrievedServiceB := container.Get((*MyServiceB)(nil)).(*MyServiceB)

	fmt.Println(retrievedServiceA.DoWorkA())
	fmt.Println(retrievedServiceB.DoWorkB())
}
