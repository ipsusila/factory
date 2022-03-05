package factory_test

import (
	"io"
	"testing"

	"github.com/ipsusila/factory"
	"github.com/stretchr/testify/assert"

	_ "github.com/ipsusila/factory/impl/file"
)

func TestFileOpenFactory(t *testing.T) {
	c := factory.ObjectConfig{
		Name: "file",
		Options: factory.Options{
			"filename": "LICENSE",
		},
	}
	fo := factory.MustCreate(c)
	fd, ok := fo.(io.ReadCloser)
	assert.True(t, ok, "Object shall implement io.ReadCloser")
	defer fd.Close()

	// read content
	buf := make([]byte, 512)
	n, err := fd.Read(buf)
	assert.Nil(t, err, "Should not error reading LICENSE file")
	assert.NotZero(t, n, "Readed bytes shall be greater than 0")

	// display

}
