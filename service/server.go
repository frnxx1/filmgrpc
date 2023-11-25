package main

import (
	"context"
	"errors"
	pb "film/service/ecommerce"
	st "film/service/storage"
	"fmt"

	"github.com/google/uuid"
)

type server struct {
	pb.UnimplementedMovieServiceServer
}

func (server) CreateMovie(ctx context.Context, req *pb.CreateFilmRequest) (*pb.CreateFilmResponse, error) {
	fmt.Println("Create Movie")
	movie := req.GetMovie()
	movie.Id = uuid.New().String()

	data := st.Film{
		ID:    movie.GetId(),
		Title: movie.GetTitle(),
		Genre: movie.GetGenre(),
	}

	res := st.DB.Create(data)
	if res.RowsAffected == 0 {
		return nil, errors.New("movie creation unsuccessful")
	}

	return &pb.CreateFilmResponse{
		Movie: &pb.Film{
			Id:    movie.GetId(),
			Title: movie.GetTitle(),
			Genre: movie.GetGenre(),
		},
	}, nil
}

func (*server) GetMovie(ctx context.Context, req *pb.ReadFilmRequest) (*pb.ReadFilmResponse, error) {
	fmt.Println("read movie")
	var movie st.Film
	res := st.DB.Find(&movie, "id = ?", req.GetId())
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}
	return &pb.ReadFilmResponse{
		Movie: &pb.Film{
			Id:    movie.ID,
			Title: movie.Title,
			Genre: movie.Genre,
		},
	}, nil
}

func (*server) GetMovies(ctx context.Context, req *pb.ReadFilmsRequest) (*pb.ReadFilmsResponse, error) {
	fmt.Println("read movies")
	movie := []*pb.Film{}
	res := st.DB.Find(&movie)
	if res.RowsAffected == 0 {
		return nil, errors.New("no movies")
	}
	return &pb.ReadFilmsResponse{Movies: movie}, nil
}

func (*server) UpdateMovie(ctx context.Context, req *pb.UpdateFilmRequest) (*pb.UpdateFilmResponse, error) {
	fmt.Println("update movie")
	var movie st.Film

	res := st.DB.Model(&movie).Where("id = ?", req.GetMovie().Id).Updates(st.Film{Title: req.Movie.Title, Genre: req.Movie.Genre})
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}

	return &pb.UpdateFilmResponse{Movie: &pb.Film{
		Id:    movie.ID,
		Title: movie.Title,
		Genre: movie.Genre,
	},
	}, nil
}

func (*server) DeleteMovie(ctx context.Context, req *pb.DeleteFilmRequest) (*pb.DeleteFilmResponse, error) {
	fmt.Println("delete film")
	var movie st.Film

	res := st.DB.Where("id = ?", req.GetId()).Delete(&movie)
	if res.RowsAffected == 0 {
		return nil, errors.New("film not found")
	}

	return &pb.DeleteFilmResponse{
		Success: true,
	}, nil
}


func DeleteMovieForTesting(ctx context.Context,title string)  {
	fmt.Println("delete film")
	var movie st.Film

	res := st.DB.Where("title = ?", title).Delete(&movie)
	if res.RowsAffected == 0 {
		fmt.Println("film not found")
	}

	
}
