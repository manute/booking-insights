package http

import (
	"booking-insights/internal/domain/stats"
	"booking-insights/internal/infrastructure/config"
	"context"
	"fmt"
	"net"
	"net/http"
)

// NewServer creates a new  http serrver with the confg passed in
func NewServer(ctx context.Context, cfg config.Config) *http.Server {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.HttPort),
		ReadTimeout:    cfg.HttpReadTimeout,
		WriteTimeout:   cfg.HttpWriteTimeout,
		MaxHeaderBytes: cfg.HttpMaxHeaderBytes,
		BaseContext: func(n net.Listener) context.Context {
			return ctx
		},
	}

	svcStats := stats.NewService()
	statsHandler := NewStatsHandler(svcStats)
	http.Handle("/stats", statsHandler)
	http.Handle("/maximize", statsHandler)

	return server
}
