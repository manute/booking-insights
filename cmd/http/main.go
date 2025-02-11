package main

import (
	"booking-insights/internal/infrastructure/config"
	htttptransport "booking-insights/internal/infrastructure/transport/http"
	"context"
	"log"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("running cmd/http ..")

	cfg, err := config.FromEnvironment()
	if err != nil {
		log.Fatalf("error loading the config: %s", err.Error())
	}

	ctx := context.TODO()
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)

	server := htttptransport.NewServer(ctx, cfg)
	defer server.Shutdown(ctx)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("http server starting...")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// running all the time until thecontext is done or an interrypt signal is received
	for {
		select {
		case <-ctx.Done():
			log.Println("ctx done, running ending")
			return
		case s := <-sigch:
			log.Printf("sgnal %s received", s)
			return
		}
	}
}
