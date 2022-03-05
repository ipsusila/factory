package factory_test

import (
	"encoding/json"
	"fmt"

	"github.com/ipsusila/factory"
)

var info = factory.Info{
	Name:        "printer",
	Description: "Example printer",
	Author:      "Your Name",
	Version:     "v0.1.0",
	Repository:  "Link to repository (if any)",
	License:     "Specify License (if any)",
}

// object to be created
type stdoutPrinter struct{}

func init() {
	factory.Register("printer", info, constructor)
}

// constructor that will be used by factory
func constructor(_ string) (factory.Object, error) {
	return &stdoutPrinter{}, nil
}

// ID implements factory.Object interface
func (p *stdoutPrinter) ID() string {
	return "StdoutPrinter"
}

// Println or other method that will be accessible through interface
func (p *stdoutPrinter) Println(args ...interface{}) {
	fmt.Println(args...)
}

func Example_implementation() {
	type printer interface {
		Println(args ...interface{})
	}

	// import package
	// import _ "path/to/package"

	// object creation configuration. `name` should be the value that was used to register factory
	configData := `{"name": "printer", "configSource": "any data to be passed to constructor"}`

	c := factory.ObjectConfig{}
	json.Unmarshal([]byte(configData), &c)

	// create object using factory
	obj, err := factory.Create(c)
	if err != nil {
		// Handler error
	}

	// convert to object
	if p, ok := obj.(printer); ok {
		p.Println("Do the real work inside object")
	}
}
