package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/CodyPerakslis/simple-grpc/messenger"
)

type server struct {
	messenger.UnimplementedMessengerServer
}

func (*server) SendMessage(ctx context.Context, mes *messenger.Message) (*messenger.Ack, error) {
	log.Println(mes.GetData())
	return &messenger.Ack{Success: true,}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	messenger.RegisterMessengerServer(grpcServer, &server{})
	log.Println("Starting server")
	grpcServer.Serve(lis)
}