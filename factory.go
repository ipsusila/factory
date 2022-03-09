package factory

import (
	"fmt"
)

// ConstructorFunc for creating object
type ConstructorFunc func(args Options) (Object, error)

// Object that will be created by the factory
type Object interface {
	ID() string
}

// Info holds factory information
type Info struct {
	Name        string
	Description string
	Version     string
	Author      string
	Repository  string
	License     string
}

// Factory that responsible for creating object
type Factory struct {
	info Info
	cf   ConstructorFunc
}

// Register factory with given information and constructor.
func Register(name string, info Info, cf ConstructorFunc) {
	f := Factory{
		info: info,
		cf:   cf,
	}
	register(name, &f)
}

// MustCreate create object using given factory name.
// If the factory does not exists, or error, it will panic
func MustCreate(c Config) Object {
	f := Get(c.Name)
	if f == nil {
		msg := fmt.Sprintf(
			"factory %s does not exist, do you forgot to import package?",
			c.Name)
		panic(msg)
	}

	obj, err := f.Create(c.Options)
	if err != nil {
		panic(err)
	}
	return obj
}

// Create create objects using given factory name and config source
func Create(c Config) (Object, error) {
	f := Get(c.Name)
	if f == nil {
		return nil, fmt.Errorf("factory %s does not exist, do you forgot to import package?",
			c.Name)
	}
	return f.Create(c.Options)
}

// Info return factory information
func (f *Factory) Info() Info {
	return f.info
}

// Create object with given configuration source
func (f *Factory) Create(args Options) (Object, error) {
	if f.cf == nil {
		return nil, fmt.Errorf("constructor is not defined in factory %s", f.info.Name)
	}
	return f.cf(args)
}
