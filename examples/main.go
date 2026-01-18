package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	unikl "github.com/Swarve-Bot/UnikL"

	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	config := &unikl.Config{
		RedisConfig: &redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		MongoConn: options.Client().ApplyURI("mongodb://localhost:10075").SetAuth(options.Credential{
			Username:   "test",
			Password:   "test",
			AuthSource: "admin",
		}),
		NatsURL: "nats://127.0.0.1:4222",
	}

	unikl, err := unikl.NewUnikl(config)
	if err != nil {
		log.Fatal(err)
	}
	defer unikl.Close()

	done := make(chan bool)

	// Suscribe
	unikl.Nats.Subscribe("example", func(msg *nats.Msg) {
		nano, _ := strconv.ParseInt(string(msg.Data), 10, 64)
		sentTime := time.Unix(0, nano)

		fmt.Printf("Latency: %v\n", time.Since(sentTime))

		done <- true
	})

	time.Sleep(100 * time.Millisecond)

	now := fmt.Sprintf("%d", time.Now().UnixNano())
	unikl.Nats.Publish("example", []byte(now))

	select {
	case <-done:
		fmt.Println("Completed.")
	case <-time.After(time.Second * 5):
		fmt.Println("Timeout reached. ")
	}
}
