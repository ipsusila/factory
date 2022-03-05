package printer

import (
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
func constructor(_ factory.Options) (factory.Object, error) {
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
