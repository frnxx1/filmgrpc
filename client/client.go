package main

import (
	"context"
	pb "film/client/ecommerce"
	st "film/service/storage"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	return r
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewMovieServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	gin.SetMode(gin.ReleaseMode)
	r := GetRouter()

	r.GET("/film/:id",func(ctx *gin.Context) {
	   id := ctx.Param("id")
       res, err := c.GetMovie(ctx, &pb.ReadFilmRequest{Id: id})
       if err != nil {
           ctx.JSON(http.StatusNotFound, gin.H{
               "message": err.Error(),
           })
           return
       }
       ctx.JSON(http.StatusOK, gin.H{
           "movie": res.Movie,
       })
	})

	r.GET("/films",func(ctx *gin.Context) {
		res,err := c.GetMovies(ctx,&pb.ReadFilmsRequest{})
		if err != nil {
			ctx.JSON(http.StatusBadRequest,gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK,res)
	})
	
	r.POST("/film", func(ct *gin.Context) {
		var film st.Film
		err := ct.ShouldBind(&film)
		if err != nil {
			ct.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data := &pb.Film{
			Title: film.Title,
			Genre: film.Genre,
		}
		res, err := c.CreateMovie(ctx, &pb.CreateFilmRequest{Movie: data})
		if err != nil {
			ct.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		ct.JSON(http.StatusCreated, gin.H{
			"film": res.Movie,
		})
	})

	r.PUT("/film/:id", func(ctx *gin.Context){
		var film st.Film
		err := ctx.ShouldBind(&film)
		if err != nil {
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err})
			return
		}
		res, err :=c.UpdateMovie(ctx,&pb.UpdateFilmRequest{
			Movie: &pb.Film{
				Id: film.ID,
				Title: film.Title,
				Genre: film.Genre,
			},
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{"updated":res})
	})

	r.DELETE("/film/:id",func(ctx *gin.Context) {
		id := ctx.Param("id")
		res, err := c.DeleteMovie(ctx,&pb.DeleteFilmRequest{Id: id})
		if err != nil {
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{"deleted":res})
	})

	r.Run("localhost:8080")

}
