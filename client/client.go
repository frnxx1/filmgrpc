package main

import (
	"context"
	pb "film/client/ecommerce"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewFilmsClient(conn)

	name := "Boys"
	description := "Boys top"
	status := "show"
	id := "1"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetFilm(ctx, &pb.FilmInfo{
		Id: id, Name: name, Description: description, Status: status,
	})
	if err != nil {
		log.Fatal("film not found", err)
	}

	log.Print("got film ", r)

}
