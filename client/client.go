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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := ClientStatus(c,ctx)
	if err != nil {
		log.Fatal("film not found ", err)
	}
	log.Print("got film ", r)

	g,err := ClientGenre(c,ctx)
	if err != nil {
		log.Fatal("film genre not found", err)
	}
	log.Print("got film ", g)
}


func ClientStatus(c pb.FilmsClient,ctx context.Context)(*pb.FilmStatus,error){
	name := "Boys"
	description := "Boys top"
	status := "show"
	id := "1"
	films := pb.FilmInfo{Id: id, Name: name, Description: description, Status: status}
	r,err := c.GetFilm(ctx,&films)
	return r, err
}

func ClientGenre(c pb.FilmsClient,ctx context.Context)(*pb.FilmGenreRole,error){
	name := "Boys"
	genre := pb.Genre_ACTION
	films := pb.FilmGenre{Name: name,Genres: &genre }
	g,err := c.GetGenre(ctx,&films)
	return g,err
}
