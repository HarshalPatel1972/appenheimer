package ratelimit

import (
	"context"
	"time"
)

type Strategy int

const (
	FailOpen Strategy = iota
	FailClosed
)

type Limiter interface {
	Allow(ctx context.Context, key string, limit int, window time.Duration) (bool, error)
}

// slidingWindowScript for Redis ZSET implementation
const slidingWindowScript = `
local key = KEYS[1]
local now = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local limit = tonumber(ARGV[3])

local clearBefore = now - window
redis.call('ZREMRANGEBYSCORE', key, 0, clearBefore)
local current = redis.call('ZCARD', key)

if current < limit then
    redis.call('ZADD', key, now, now)
    redis.call('EXPIRE', key, math.ceil(window/1000))
    return 1
else
    return 0
end
`
