package unikl

import (
	"context"

	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Unikl struct {
	MongoConn *mongo.Client

	Nats      *nats.Conn
	Jetstream nats.JetStreamContext
	MongoX    *mongox.Client
	Redis     *redis.Client
}

func NewUnikl(config *Config) (*Unikl, error) {

	nc, err := nats.Connect(config.NatsURL)
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	redisConn := redis.NewClient(config.RedisConfig)

	mongoConn, err := mongo.Connect(config.MongoConn)
	if err != nil {
		return nil, err
	}

	mongo := mongox.NewClient(mongoConn, &mongox.Config{})

	return &Unikl{
		Redis:     redisConn,
		MongoConn: mongoConn,

		Nats:      nc,
		Jetstream: js,
		MongoX:    mongo,
	}, nil
}

func (u *Unikl) Close() error {
	if err := u.Redis.Close(); err != nil {
		return err
	}
	if err := u.MongoX.Disconnect(context.Background()); err != nil {
		return err
	}
	if err := u.MongoConn.Disconnect(context.Background()); err != nil {
		return err
	}
	u.Nats.Close()

	return nil
}
