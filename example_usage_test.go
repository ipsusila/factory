package factory_test

import (
	"io"
	"log"

	"github.com/ipsusila/factory"

	_ "github.com/ipsusila/factory/impl/file"
)

func Example_usage() {
	// Create file reader object using registered factory
	fo := factory.MustCreate(factory.ObjectConfig{"file", "LICENSE"})
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
