package file

import (
	"os"

	"github.com/ipsusila/factory"
)

type fileOpener struct {
	*os.File
}

var info = factory.Info{
	Name:        "file",
	Description: "Example file opener object. Don't forget to close after using it.",
	Author:      "I Putu Susila",
	Version:     "v0.1.0",
	Repository:  "github.com/ipsusila/factory/file",
	License:     "MIT",
}

func init() {
	// Register factory
	factory.Register("file", info, constructor)
}

// constructor create fileObject that open file from given configSrc
func constructor(configSrc string) (factory.Object, error) {
	fd, err := os.Open(configSrc)
	if err != nil {
		return nil, err
	}
	return &fileOpener{fd}, nil
}

// ID return object identifier
func (f *fileOpener) ID() string {
	return "FileOpener"
}
