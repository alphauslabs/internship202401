package main

import (
	"context"
	"flag"
	"log"

	"cloud.google.com/go/pubsub"
)

var (
	topic = flag.String("topic", "", "Topic name for publishing")
	msg   = flag.String("msg", "", "Message to publish")
)

func main() {
	projectId := "mobingi-main"
	ctx := context.Background()

	if *topic == "" {
		log.Println("topic cannot be empty")
		return
	}

	if *msg == "" {
		log.Println("msg cannot be empty")
		return
	}

	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Println("NewClient failed:", err)
		return
	}

	defer client.Close()
	t := client.Topic(*topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(*msg),
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		log.Println("Get failed:", err)
		return
	}

	log.Printf("Published %q; id=%v\n", *msg, id)
}
