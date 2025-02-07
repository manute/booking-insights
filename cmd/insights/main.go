package main

import (
	"context"
	"log"
	"os"
	"time"
)

const defaultTimeout = time.Minute

func main() {
	timeout := defaultTimeout
	var err error

	outVal := os.Getenv("TIMEOUT")
	if len(outVal) > 0 {
		timeout, err = time.ParseDuration(outVal)
		if err != nil {
			log.Fatalf("Error parsing timeout duration %v", err)
		}

	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			log.Println("Ended at:", time.Now())
			return
		case <-time.Tick(1 * time.Second):
			log.Println(".")
		}
	}

}
