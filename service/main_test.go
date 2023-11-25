package main

import (
	"context"
	pb "film/service/ecommerce"
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
func createTest(ctx context.Context) *pb.Film {
	data := &pb.Film{
		Title: "Testing and Delete",
		Genre: "Testing and Delete",
	}
	res, _ := server.CreateMovie(server{}, ctx, &pb.CreateFilmRequest{Movie: data})

	return res.Movie
}

func TestCreateMovie(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)

	data := createTest(ctx)

	resp, err := client.CreateMovie(ctx, &pb.CreateFilmRequest{Movie: data})
	if err != nil {
		t.Fatalf("CreateMovie failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

	DeleteMovieForTesting(context.Background(), data.Title)
}

func TestReadMovie(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)
	data := createTest(ctx)

	resp, err := client.GetMovie(ctx, &pb.ReadFilmRequest{Id: data.Id})
	if err != nil {
		t.Fatalf("GetMovie failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

	DeleteMovieForTesting(context.Background(), data.Title)
}

func TestReadMovies(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)
	data := createTest(ctx)
	resp, err := client.GetMovies(ctx, &pb.ReadFilmsRequest{})
	if err != nil {
		t.Fatalf("GetMovies failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

	DeleteMovieForTesting(context.Background(), data.Title)
}

func TestUpdateMovies(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)

	data := createTest(ctx)

	resp, err := client.UpdateMovie(ctx, &pb.UpdateFilmRequest{
		Movie: &pb.Film{
			Id:    data.Id,
			Title: data.Title,
			Genre: data.Genre,
		},
	})
	if err != nil {
		t.Fatalf("UpdateMovies failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

	DeleteMovieForTesting(context.Background(), data.Title)

}

func TestDeleteMovies(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)

	data := createTest(ctx)

	resp, err := client.DeleteMovie(ctx, &pb.DeleteFilmRequest{Id: data.Id})
	if err != nil {
		t.Fatalf("UpdateMovies failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

	DeleteMovieForTesting(context.Background(), data.Title)
}
