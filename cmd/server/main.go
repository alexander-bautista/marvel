package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	rest "github.com/alexander-bautista/marvel/pkg/api"
	"github.com/alexander-bautista/marvel/pkg/comic"
	"github.com/alexander-bautista/marvel/pkg/repository/mongo"
)

func main() {
	repo := chooseRepo()
	service := comic.NewComicService(repo)
	handler := rest.NewHandler(service)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Error while shutdown server:", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	//Disconnect(ctx, col)
	log.Println("Server exiting")

}

func chooseRepo() comic.ComicRepository {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		mongoURL = "mongodb+srv://todo_user:todo2020@traffic-nkwxe.mongodb.net/todo?retryWrites=true&w=majority"
	}
	mongodb := os.Getenv("MONGO_DB")
	if mongodb == "" {
		mongodb = "todo"
	}

	mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))

	if mongoTimeout == 0 {
		mongoTimeout = 10
	}

	repo, err := mongo.NewMongoRepository(mongoURL, mongodb, mongoTimeout)

	if err != nil {
		log.Fatal(err)
	}

	return repo

}
