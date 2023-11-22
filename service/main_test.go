package main

import (
	"context"
	pb "film/service/ecommerce"
	st "film/service/storage"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterMovieServiceServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateMovie(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)

	var film st.Film

	data := &pb.Film{
		Title: film.Title,
		Genre: film.Genre,
	}

	resp, err := client.CreateMovie(ctx, &pb.CreateFilmRequest{Movie: data})
	if err != nil {
		t.Fatalf("CreateMovie failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	
}

func TestReadMovie(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)

	resp, err := client.GetMovie(ctx, &pb.ReadFilmRequest{Id: "c5900158-7e9f-497d-9595-24d24a6ee480"})
	if err != nil {
		t.Fatalf("GetMovie failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	
}

func TestReadMovies(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)

	resp, err := client.GetMovies(ctx, &pb.ReadFilmsRequest{})
	if err != nil {
		t.Fatalf("GetMovies failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	
}

func TestUpdateMovies(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)


	resp, err := client.UpdateMovie(ctx, &pb.UpdateFilmRequest{
		Movie: &pb.Film{
			Id:    "c5900158-7e9f-497d-9595-24d24a6ee480",
			Title: "film.Title",
			Genre: "film.Genre",
		},
	})
	if err != nil {
		t.Fatalf("UpdateMovies failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	
}

func TestDeleteMovies(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)


	resp, err := client.DeleteMovie(ctx, &pb.DeleteFilmRequest{Id: "4c9a4215-443a-4abf-aaa2-edf8c452cb85"})
	if err != nil {
		t.Fatalf("UpdateMovies failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	
}
