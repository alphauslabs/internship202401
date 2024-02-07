package main

import (
	"context"
	"flag"
	"log"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	subId = flag.String("subscription", "", "Subscription name")
)

func main() {
	flag.Parse()
	projectId := "mobingi-main"
	ctx := context.Background()

	if *subId == "" {
		log.Println("subscription cannot be empty")
		return

	}

	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Println("NewClient failed:", err)
		return
	}

	defer client.Close()
	sub := client.Subscription(*subId)
	sub.ReceiveSettings.Synchronous = true

	// Receive messages for 10 seconds.
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var received int32

	// Receive blocks until the context is cancelled or an error occurs.
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		atomic.AddInt32(&received, 1)
		log.Printf("Received: %q", msg)
		msg.Ack()
	})

	if err != nil {
		log.Println("Receive failed:", err)
		return
	}

	log.Printf("Received %d messages\n", received)
}
