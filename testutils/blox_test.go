package testutils

import (
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"
)

func TestBloxTransportImplements(t *testing.T) {
	b := new(BloxTransport)
	assert.Implements(t, (*runtime.ClientTransport)(nil), b)
}
