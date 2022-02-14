package redis_test

import (
	"testing"

	"actions/internal/client/redis/redistest"

	"github.com/stretchr/testify/assert"
)

func Test_Integration_Redis_Client_Prints_Msg(t *testing.T) {
	// arrange
	SUT := redistest.Redis(t)

	// act
	output := SUT.HelloWorld()

	// assert
	assert.Equal(t, "hello (integration testing) world", output)
}
