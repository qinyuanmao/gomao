package security

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRSA(t *testing.T) {
	var parser = NewDefaultParser()
	var input = []byte("hello world")
	output, err := parser.Encode(input)
	assert.Nil(t, err)
	decode, err := parser.Decode(output)
	assert.Nil(t, err)
	assert.Equal(t, input, decode)
	os.Remove("./config")
}
