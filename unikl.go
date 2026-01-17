package unikl

import (
	"context"

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

	mongoConn, err := mongo.Connect(config.MongoConn)
	if err != nil {
		return nil, err
	}

	mongo := mongox.NewClient(mongoConn, &mongox.Config{})

	return &Unikl{
		RedisConn: redisConn,
		MongoConn: mongoConn,

		RabbitMQ: rbmq,
		MongoX:   mongo,
		Cache:    cache,
	}, nil
}

func (u *Unikl) Close() error {
	if err := u.RedisConn.Close(); err != nil {
		return err
	}
	if err := u.MongoX.Disconnect(context.Background()); err != nil {
		return err
	}
	if err := u.MongoConn.Disconnect(context.Background()); err != nil {
		return err
	}
	if err := u.RabbitMQ.Close(); err != nil {
		return err
	}

	return nil
}
