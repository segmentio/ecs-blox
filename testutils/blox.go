package testutils

import (
	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/mock"
)

// BloxTransport is a mock-able ClientTransport that can be used in a fake blox.
type BloxTransport struct {
	mock.Mock
}

// Submit knows how to complete an API call.
func (m *BloxTransport) Submit(op *runtime.ClientOperation) (interface{}, error) {
	args := m.Called(op)
	return args.Get(0), args.Error(1)
}
