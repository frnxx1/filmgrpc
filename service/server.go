package main

import (
	"context"

	pb "film/service/ecommerce"
)

type server struct {
	*pb.FilmStatus
	pb.UnimplementedFilmsServer
}

func (s *server) GetFilm(ctx context.Context, in *pb.FilmInfo) (*pb.FilmStatus, error) {
	return &pb.FilmStatus{
		Status: in.Status,
		}, nil
}
