package unikl

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	NatsURL     string
	RedisConfig *redis.Options
	MongoConn   *options.ClientOptions
}
