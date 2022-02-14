package redis

import "errors"

// RedisHostEnvVar is the name of the env var used to set where the app looks for Redis
const RedisHostEnvVar = "REDIS_HOST"

// RedisClient is a mock client wrapper for Redis.
// It will serve the demonstration because it is typically a
// managed out-of-process dependency and therefore a candidate
// for an integration test
type RedisClient struct{}

func NewRedisClient(host string) (*RedisClient, error) {
	if host == "" {
		return nil, errors.New("redis host is empty string")
	}

	// this would normally connect to the redis instance
	// in a non-demonstrational pkg
	return &RedisClient{}, nil
}

func (r *RedisClient) HelloWorld() string {
	return "hello (integration testing) world"
}
