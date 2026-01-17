package unikl

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Config struct {
	RabbitMQURL string
	RedisConfig *redis.Options
	MongoConn   *mongo.Client
}
