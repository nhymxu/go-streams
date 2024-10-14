package redisstream

import "github.com/nhymxu/go-streams/extensions"

type RedisConsumer struct {
}

// RedisConsumer must satisfy the IConsumer interface
var _ extensions.IConsumer = (*RedisConsumer)(nil)

func (c RedisConsumer) Consume() error {
	return nil
}
