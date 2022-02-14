package redistest

import (
	"os"
	"testing"

	"actions/internal/client/redis"

	"github.com/stretchr/testify/require"
)

// RedisIntegrationTestHostEnvVar is where the test suite looks for the Redis host
const RedisIntegrationTestHostEnvVar = "REDIS_INTEGRATION_TEST_HOST"

func Redis(t *testing.T) *redis.RedisClient {
	host := os.Getenv(RedisIntegrationTestHostEnvVar)
	if host == "" {
		t.Skip()
	}

	client, err := redis.NewRedisClient(host)
	require.Nil(t, err)

	return client
}
