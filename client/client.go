package main

import (
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	"github.com/CodyPerakslis/simple-grpc/messenger"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("simplegrpc-server:8084", opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := messenger.NewMessengerClient(conn)

	message := &messenger.Message{Data: "This is a test message.\n",}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ack, err := client.SendMessage(ctx, message)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println(ack)
}