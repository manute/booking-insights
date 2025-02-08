package main

import (
	"context"
	"log"
	"os"
	"time"
)

func main() {
	var timeout time.Duration = time.Minute
	var err error

	if val, ok := os.LookupEnv("TIMEOUT"); ok && len(val) > 0 {
		timeout, err = time.ParseDuration(val)
		if err != nil {
			log.Fatalf("error parsing timeout: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			log.Printf("running ernded at: %s", time.Now())
			return
		case <-time.Tick(1 * time.Second):
			log.Println(".")
		}
	}
}
