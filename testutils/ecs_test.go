package testutils

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/stretchr/testify/assert"
)

func TestECSAPIImplemented(t *testing.T) {
	e := new(ECSAPI)
	assert.Implements(t, (*ecsiface.ECSAPI)(nil), e)
}
