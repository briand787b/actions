package redis_test

import (
	"testing"

	"actions/internal/client/redis/redistest"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationRedisClientPrintsMsg(t *testing.T) {
	// arrange
	SUT := redistest.Redis(t)

	// act
	output := SUT.HelloWorld()

	// assert
	assert.Equal(t, "hello (integration testing) world", output)
}
