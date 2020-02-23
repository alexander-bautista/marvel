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
	"github.com/alexander-bautista/marvel/pkg/character"
	"github.com/alexander-bautista/marvel/pkg/comic"
	"github.com/alexander-bautista/marvel/pkg/repository/mongo"
)

func main() {
	mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))

	if mongoTimeout == 0 {
		mongoTimeout = 10
	}

	client, err := mongo.NewMongoClient()

	defer client.Disconnect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	comicRepo, err := mongo.NewMongoComicRepository(mongoTimeout, client)

	if err != nil {
		log.Fatal(err)
	}

	characterRepo, err := mongo.NewMongoCharacterRepository(mongoTimeout, client)

	if err != nil {
		log.Fatal(err)
	}
	comicService := comic.NewComicService(comicRepo)
	characterService := character.NewCharacterService(characterRepo)
	handler := rest.NewHandler(comicService, characterService)

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
