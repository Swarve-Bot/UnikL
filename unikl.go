package unikl

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/redis/go-redis/v9"
	"github.com/wagslane/go-rabbitmq"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/eko/gocache/lib/v4/cache"
	redis_store "github.com/eko/gocache/store/redis/v4"
)

type Unikl struct {
	RedisConn *redis.Client
	MongoConn *mongo.Client

	RabbitMQ *rabbitmq.Conn
	MongoX   *mongox.Client
	Cache    *cache.Cache[string]
}

func NewUnikl(config *Config) (*Unikl, error) {
	rbmq, err := rabbitmq.NewConn(config.RabbitMQURL, rabbitmq.WithConnectionOptionsLogging)
	if err != nil {
		return nil, err
	}

	redisConn := redis.NewClient(config.RedisConfig)

	store := redis_store.NewRedis(redisConn)
	cache := cache.New[string](store)

	mongo := mongox.NewClient(config.MongoConn, &mongox.Config{})

	return &Unikl{
		RedisConn: redisConn,
		MongoConn: config.MongoConn,

		RabbitMQ: rbmq,
		MongoX:   mongo,
		Cache:    cache,
	}, nil
}
