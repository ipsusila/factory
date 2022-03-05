package factory_test

import (
	"io"
	"log"

	"github.com/ipsusila/factory"

	_ "github.com/ipsusila/factory/impl/file"
)

func Example_usage() {
	conf := factory.ObjectConfig{
		Name: "file",
		Options: factory.Options{
			"filename": "LICENSE",
		},
	}
	// Create file reader object using registered factory
	fo := factory.MustCreate(conf)
	fd, ok := fo.(io.ReadCloser)
	if !ok {
		log.Fatalf("FileReader object must implement io.ReadCloser interface")
	}
	defer fd.Close()

	// Read content using created file reader object
	buf := make([]byte, 512)
	_, err := fd.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
}
