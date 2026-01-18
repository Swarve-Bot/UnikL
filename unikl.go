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
	var mongoConn *mongo.Client
	var mongoX *mongox.Client
	var redisConn *redis.Client
	var nc *nats.Conn
	var js nats.JetStreamContext
	var err error

	if config.UseNats {
		nc, err := nats.Connect(config.NatsURL)
		if err != nil {
			return nil, err
		}

		js, err = nc.JetStream()
		if err != nil {
			return nil, err
		}
	}

	if config.UseRedis {
		redisConn = redis.NewClient(config.RedisConfig)
	}

	if config.UseMongo {
		mongoConn, err = mongo.Connect(config.MongoConn)
		if err != nil {
			return nil, err
		}
		mongoX = mongox.NewClient(mongoConn, &mongox.Config{})
	}

	return &Unikl{
		Redis:     redisConn,
		MongoConn: mongoConn,

		Nats:      nc,
		Jetstream: js,
		MongoX:    mongoX,
	}, nil
}

func (u *Unikl) Close() error {
	if u.Nats != nil {
		u.Nats.Close()
	}

	if u.Redis != nil {
		if err := u.Redis.Close(); err != nil {
			return err
		}
	}

	if u.MongoX != nil {
		if err := u.MongoX.Disconnect(context.Background()); err != nil {
			return err
		}
	}

	if u.MongoConn != nil {
		if err := u.MongoConn.Disconnect(context.Background()); err != nil {
			return err
		}
	}

	return nil
}
