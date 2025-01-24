package main

import (
	"context"
	"fmt"
	http "github.com/krisnaadi/dashboard-cronjob-be/internal/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	server := http.NewHTTP(ctx)

	go func() {
		err := server.Run().ListenAndServe()
		if err != nil {
			log.Fatal("server.Run().ListenAndServe() error - mainHTTP")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	fmt.Printf("Server stopped\n")

	os.Exit(0)
}
