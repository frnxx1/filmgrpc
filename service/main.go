package main

import (
	pb "film/service/ecommerce"
	"flag"
	"fmt"
	"log"
	"net"
	"film/service/metrics"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "gRPC server port")
 )

func main() {
	lis, err := net.Listen("tcp",fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterMovieServiceServer(s, &server{})
	go func() {
		_ = metrics.Listen("127.0.0.1:8082")
	}()

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
