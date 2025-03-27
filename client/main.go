package main

import (
	"context"
	"github.com/bench/kingdom-server/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetUser(ctx, &pb.UserRequest{UserId: "123"})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	log.Printf("User Info:\nName: %s\nAge: %d\nEmail: %s",
		res.Name, res.Age, res.Email)
}
