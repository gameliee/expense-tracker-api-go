package tools

import (
	"fmt"
	"reflect"
)

type DependencyInjector interface {
	RegisterInstance(interface{})
	Get(instanceType interface{}) interface{}
	Build() error
}

type Container struct {
	instances map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		instances: make(map[string]interface{}),
	}
}

func (c *Container) RegisterInstance(instance interface{}) {
	typ := reflect.TypeOf(instance)
	c.instances[typ.String()] = instance
}

// Get returns the instance by type
func (c *Container) Get(instanceType interface{}) interface{} {
	// Convert the type of the instance to its string representation
	typ := reflect.TypeOf(instanceType)
	if instance, ok := c.instances[typ.String()]; ok {
		return instance
	}
	return nil
}

// Build automatically injects dependencies using reflection and struct tags
// If any field cannot be injected after all tries, it raises an error
func (c *Container) Build() error {
	for try := 0; try < 10; try++ { // Try to ensure all dependencies are injected
		for _, instance := range c.instances {
			val := reflect.ValueOf(instance).Elem()
			typElem := val.Type()

			for i := 0; i < val.NumField(); i++ {
				field := val.Field(i)
				fieldType := field.Type()
				fieldTag := typElem.Field(i).Tag.Get("inject")

				// If the field has the "inject" tag, we attempt to inject
				if fieldTag != "" && field.CanSet() && field.IsZero() {
					// Look up the dependency in the container by type
					if depInstance, ok := c.instances[fieldTag]; ok && reflect.TypeOf(depInstance).AssignableTo(fieldType) {
						field.Set(reflect.ValueOf(depInstance))
					}
				}
			}
		}
	}
	// After all tries, check if there are any unresolved (nil) fields
	for _, instance := range c.instances {
		val := reflect.ValueOf(instance).Elem()
		typElem := val.Type()

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldTag := typElem.Field(i).Tag.Get("inject")

			// If the field has the "inject" tag and is still nil
			if fieldTag != "" && field.CanSet() && field.IsZero() {
				return fmt.Errorf("failed to inject dependency for %s.%s", typElem.Name(), typElem.Field(i).Name)
			}
		}
	}
	return nil
}
