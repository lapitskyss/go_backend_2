package greeter

import (
	"context"
	"log"
)

func SayHello(_ context.Context, name string) string {
	log.Printf("Received: %s", name)

	return "Hello " + name
}
