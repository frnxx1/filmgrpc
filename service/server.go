package main

import (
	"context"

	pb "film/service/ecommerce"
)

type server struct {
	*pb.FilmStatus
	pb.UnimplementedFilmsServer
	*pb.FilmGenreRole
	*pb.FilmUpdate
}

func (s *server) GetFilm(ctx context.Context, in *pb.FilmInfo) (*pb.FilmStatus, error) {
	return &pb.FilmStatus{
		Status: in.Status,
	}, nil
}

func (s *server) GetGenre(ctx context.Context, in *pb.FilmGenre) (*pb.FilmGenreRole, error) {
	value := &pb.FilmGenreRole{Genres: in.Genres}
	return &pb.FilmGenreRole{Genres: value.Genres.Enum()}, nil
}

func (s *server) UpdateStatus(ctx context.Context, in *pb.FilmInfo)(*pb.FilmUpdate,error){
	value := &pb.FilmInfo{Status: in.Status}
	if value.Status{
		return &pb.FilmUpdate{NewStatus: value.Status}, nil
	}
	return &pb.FilmUpdate{NewStatus: true}, nil
}
