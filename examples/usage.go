package examples

import (
	unikl "github.com/Swarve-Bot/UnikL"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	// Initialize Unikl instance
	config := &unikl.Config{
		RedisConfig: &redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		MongoConn: options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
			Username:   "test",
			Password:   "test",
			AuthSource: "db-test",
		}),
		RabbitMQURL: "amqp://localhost:5672",
	}

	unikl, err := unikl.NewUnikl(config)
	if err != nil {
		panic(err)
	}

	// Use Unikl instance
	// ...

	// Close Unikl instance
	unikl.Close()
}
