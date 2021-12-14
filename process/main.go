package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/mediocregopher/radix/v3"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/kafkapubsub"
)

var (
	pool     *radix.Pool
	sub      *pubsub.Subscription
	connFunc = func(network, addr string) (radix.Conn, error) {
		return radix.Dial(network, addr,
			radix.DialTimeout(1*time.Minute),
		)
	}
)

func main() {
	for {
		s, err := subscription()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}
		msg, err := s.Receive(context.Background())
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}

		log.Printf("Reseved message: %s\n", string(msg.Body))

		err = storage().Do(radix.Cmd(nil, "LPUSH", "result", string(msg.Body)))
		if err != nil {
			log.Println(err)
		}
		if rand.Float64() < .05 {
			_ = storage().Do(radix.Cmd(nil, "LTRIM", "result", "0", "9"))
		}
		msg.Ack()
	}
}

func subscription() (*pubsub.Subscription, error) {
	if sub != nil {
		return sub, nil
	}
	var err error
	sub, err = pubsub.OpenSubscription(context.Background(), "kafka://process?topic=rates")
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func storage() *radix.Pool {
	if pool != nil {
		return pool
	}
	var err error
	pool, err = radix.NewPool("tcp", "redis:6379", 1, radix.PoolConnFunc(connFunc))
	if err != nil {
		panic(err)
	}
	return pool
}
